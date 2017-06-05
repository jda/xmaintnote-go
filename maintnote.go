package xmaintnote

//
// MaintNote fields
//
const maintProvider string = "X-MAINTNOTE-PROVIDER"
const maintAccount string = "X-MAINTNOTE-ACCOUNT"
const maintMaintID string = "X-MAINTNOTE-MAINTENANCE-ID"
const maintObjectID string = "X-MAINTNOTE-OBJECT-ID"
const maintImpact string = "X-MAINTNOTE-IMPACT"
const maintStatus string = "X-MAINTNOTE-STATUS"

//
// Impacts
//

// ImpactNone represents MAINTNOTE NO-IMPACT impact
const ImpactNone string = "NO-IMPACT"

// ImpactReducedRedundancy represents the MAINTNOTE REDUCED-REDUNDANCY impact
const ImpactReducedRedundancy string = "REDUCED-REDUNDANCY"

// ImpactDegraded represents the MAINTNOTE DEGRADED impact
const ImpactDegraded string = "DEGRADED"

// ImpactOutage represents the MAINTNOTE OUTAGE impact
const ImpactOutage string = "OUTAGE"

//
// Statuses
//

// StatusTenative represents the MAINTNOTE TENATIVE status
const StatusTenative string = "TENATIVE"

// StatusCancelled represents the MAINTNOTE CANCELLED status
const StatusCancelled string = "CANCELLED"

// StatusInProcess represents the MAINTNOTE IN-PROCESS status
const StatusInProcess string = "IN-PROCESS"

// StatusCompleted represents the MAINTNOTE COMPLETED status
const StatusCompleted string = "COMPLETED"

// ValidStatus checks if a status value is a valid X-MAINTNOTE status
func ValidStatus(status string) bool {
	if status == StatusTenative {
		return true
	}

	if status == StatusCancelled {
		return true
	}

	if status == StatusInProcess {
		return true
	}

	if status == StatusCompleted {
		return true
	}

	return false
}

// ValidImpact checks if a status value is a valid X-MAINTNOTE impact
func ValidImpact(impact string) bool {
	if impact == ImpactNone {
		return true
	}

	if impact == ImpactReducedRedundancy {
		return true
	}

	if impact == ImpactDegraded {
		return true
	}

	if impact == ImpactOutage {
		return true
	}

	return false
}
