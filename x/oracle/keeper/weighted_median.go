package keeper

import (
	"context"
	"errors"
	"math/big"
	"sort"

	"github.com/tellor-io/layer/x/oracle/types"

	cosmomath "cosmossdk.io/math"
)

func (k Keeper) WeightedMedian(ctx context.Context, reports []types.MicroReport) (*types.Aggregate, error) {
	var medianReport types.Aggregate
	values := make(map[string]cosmomath.LegacyDec)

	for _, r := range reports {
		val, ok := new(big.Int).SetString(r.Value, 16)
		if !ok {
			k.Logger(ctx).Error("WeightedMedian", "error", "failed to parse value")
			return nil, errors.New("failed to parse value")
		}
		values[r.Reporter] = cosmomath.LegacyNewDecFromBigInt(val)
	}

	sort.SliceStable(reports, func(i, j int) bool {
		return values[reports[i].Reporter].BigInt().Cmp(values[reports[j].Reporter].BigInt()) < 0
	})

	totalReporterPower, weightedSum := cosmomath.LegacyZeroDec(), cosmomath.LegacyZeroDec()
	for _, r := range reports {
		weightedSum = weightedSum.Add(values[r.Reporter].Mul(cosmomath.LegacyNewDec(r.Power)))
		totalReporterPower = totalReporterPower.Add(cosmomath.LegacyNewDec(r.Power))
		medianReport.Reporters = append(medianReport.Reporters, &types.AggregateReporter{Reporter: r.Reporter, Power: r.Power, BlockNumber: r.BlockNumber})
	}

	halfTotalPower := totalReporterPower.Quo(cosmomath.LegacyNewDec(2))
	cumulativePower := cosmomath.LegacyZeroDec()

	// Find the weighted median
	for i, s := range reports {
		cumulativePower = cumulativePower.Add(cosmomath.LegacyNewDec(s.Power))
		if cumulativePower.BigInt().Cmp(halfTotalPower.BigInt()) >= 0 {
			medianReport.ReporterPower = totalReporterPower.TruncateInt64()
			medianReport.AggregateReporter = s.Reporter
			medianReport.AggregateValue = s.Value
			medianReport.QueryId = s.QueryId
			medianReport.AggregateReportIndex = int64(i)
			medianReport.MicroHeight = s.BlockNumber
			break
		}
	}

	// Calculate the weighted standard deviation
	sumWeightedSquaredDiffs := cosmomath.LegacyZeroDec()
	weightedMeanDec := weightedSum.Quo(totalReporterPower)

	for _, r := range reports {
		diffDec := values[r.Reporter].Sub(weightedMeanDec)
		diffDecSquared := diffDec.Mul(diffDec)
		weightedSquaredDiffDec := diffDecSquared.Mul(cosmomath.LegacyNewDec(r.Power))
		sumWeightedSquaredDiffs = sumWeightedSquaredDiffs.Add(weightedSquaredDiffDec)
	}

	weightedVariance := sumWeightedSquaredDiffs.Quo(totalReporterPower)
	wstdDeviation, err := weightedVariance.ApproxSqrt()
	if err != nil {
		k.Logger(ctx).Error("WeightedMedian", "error", "failed to calculate weighted standard deviation")
		medianReport.StandardDeviation = "0"
	} else {
		medianReport.StandardDeviation = wstdDeviation.String()
	}

	err = k.SetAggregate(ctx, &medianReport)
	if err != nil {
		return nil, err
	}
	return &medianReport, nil
}
