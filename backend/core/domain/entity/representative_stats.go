package entity

// RepresentativeStats represents statistics for a representative
type RepresentativeStats struct {
	uuid               string
	trainingCount      int64
	onlineActionCount  int64
	offlineActionCount int64
	offlineActionValue int64
	showroomItemCount  int64
	repMarketingCount  int64
}

// NewRepresentativeStats creates a new RepresentativeStats
func NewRepresentativeStats(
	uuid string,
	trainingCount int64,
	onlineActionCount int64,
	offlineActionCount int64,
	offlineActionValue int64,
	showroomItemCount int64,
	repMarketingCount int64,
) *RepresentativeStats {
	return &RepresentativeStats{
		uuid:               uuid,
		trainingCount:      trainingCount,
		onlineActionCount:  onlineActionCount,
		offlineActionCount: offlineActionCount,
		offlineActionValue: offlineActionValue,
		showroomItemCount:  showroomItemCount,
		repMarketingCount:  repMarketingCount,
	}
}

// Getters

func (rs *RepresentativeStats) UUID() string {
	return rs.uuid
}

func (rs *RepresentativeStats) OnlineActionCount() int64 {
	return rs.onlineActionCount
}

func (rs *RepresentativeStats) OfflineActionCount() int64 {
	return rs.offlineActionCount
}

func (rs *RepresentativeStats) OfflineActionValue() int64 {
	return rs.offlineActionValue
}

func (rs *RepresentativeStats) ShowroomItemCount() int64 {
	return rs.showroomItemCount
}

func (rs *RepresentativeStats) RepMarketingCount() int64 {
	return rs.repMarketingCount
}

func (rs *RepresentativeStats) TrainingCount() int64 {
	return rs.trainingCount
}

// TotalActions returns the total number of actions (online + offline)
func (rs *RepresentativeStats) TotalActions() int64 {
	return rs.onlineActionCount + rs.offlineActionCount
}
