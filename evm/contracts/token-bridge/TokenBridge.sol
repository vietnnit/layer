// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import "../bridge/BlobstreamO.sol";
import "../interfaces/IERC20.sol";

contract TokenBridge{
    BlobstreamO public bridge;
    IERC20 public token;
    uint256 public currentDepositLimit;
    uint256 public depositId;
    uint256 public depositLimitUpdateTime;
    uint256 public constant DEPOSIT_LIMIT_UPDATE_INTERVAL = 12 hours;
    uint256 public constant INITIAL_LAYER_TOKEN_SUPPLY = 100 ether; // update this as needed
    uint256 public constant MAX_ATTESTATION_AGE = 12 hours;
    uint256 public immutable DEPOSIT_LIMIT_DENOMINATOR = 100e18 / 20e18; // 100/depositLimitPercentage

    mapping(uint256 => bool) public withdrawalClaimed;
    mapping(uint256 => DepositDetails) public deposits;

    struct DepositDetails {
        address sender;
        string recipient;
        uint256 amount;
        uint256 blockHeight;
    }

    event Deposit(uint256 _depositId, address _sender, string _recipient, uint256 _amount);
    event Withdrawal(uint256 _depositId, string _sender, address _recipient, uint256 _amount);

    constructor(address _token, address _blobstream){
        token = IERC20(_token);
        bridge = BlobstreamO(_blobstream);
        _depositLimit();
    }

    function depositToLayer(uint256 _amount, string memory _layerRecipient) external {
        require(_amount > 0, "TokenBridge: amount must be greater than 0");
        require(_amount <= _depositLimit(), "TokenBridge: amount exceeds deposit limit");
        require(token.transferFrom(msg.sender, address(this), _amount), "TokenBridge: transferFrom failed");
        depositId++;
        currentDepositLimit -= _amount;
        deposits[depositId] = DepositDetails(msg.sender, _layerRecipient, _amount, block.number);
        emit Deposit(depositId, msg.sender, _layerRecipient, _amount);
    }

    function withdrawFromLayer(
        OracleAttestationData calldata _attest,
        Validator[] calldata _valset,
        Signature[] calldata _sigs,
        uint256 _depositId
    ) external {
        require(_attest.queryId == keccak256(abi.encode("TRBBridge", abi.encode(false, _depositId))), "TokenBridge: invalid queryId");
        require(!withdrawalClaimed[_depositId], "TokenBridge: withdrawal already claimed");
        require(block.timestamp - _attest.report.timestamp > 12 hours, "TokenBridge: premature attestation");
        //isAnyConsesnusValue here
        require(bridge.verifyConsensusOracleData(_attest, _valset, _sigs), "Invalid attestation");
        require(block.timestamp - _attest.attestationTimestamp <= MAX_ATTESTATION_AGE , "Attestation is too old");
        //to here
        withdrawalClaimed[_depositId] = true;    
        (address _recipient, string memory _layerSender,uint256 _amountLoya) = abi.decode(_attest.report.value, (address, string, uint256));
        uint256 _amountConverted = _amountLoya * 1e12; 
        require(token.transfer(_recipient, _amountConverted), "TokenBridge: transfer failed");
        emit Withdrawal(_depositId, _layerSender, _recipient, _amountConverted);
    }

    function _depositLimit() internal returns (uint256) {
        if (block.timestamp - depositLimitUpdateTime > DEPOSIT_LIMIT_UPDATE_INTERVAL) {
            uint256 _layerTokenSupply = token.balanceOf(address(this)) + INITIAL_LAYER_TOKEN_SUPPLY;
            currentDepositLimit = _layerTokenSupply / DEPOSIT_LIMIT_DENOMINATOR;
            depositLimitUpdateTime = block.timestamp;
        }
        return currentDepositLimit;
    }
}
