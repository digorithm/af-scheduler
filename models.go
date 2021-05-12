package main

type TimeSlot struct {
	ClassName        string `json:"class_name"`
	ID               string `json:"id"`
	ClubID           string `json:"club_id"`
	ClubIdent        string `json:"club_ident"`
	CompanyIdent     string `json:"company_ident"`
	Date             string `json:"date"`
	StartTime        string `json:"start_time"`
	StopTime         string `json:"stop_time"`
	Timezone         string `json:"timezone"`
	StartInt         int64  `json:"start_int"`
	StopInt          int64  `json:"stop_int"`
	TotalMinutes     int    `json:"total_minutes"`
	WorkoutMinutes   int    `json:"workout_minutes"`
	CleaningMinutes  int    `json:"cleaning_minutes"`
	TotalSpots       int    `json:"total_spots"`
	StaffSpots       int    `json:"staff_spots"`
	WalkinSpots      int    `json:"walkin_spots"`
	AdvanceSpots     int    `json:"advance_spots"`
	IsSenior         bool   `json:"is_senior"`
	CreatedInt       int64  `json:"created_int"`
	UpdatedInt       int64  `json:"updated_int"`
	AdvanceSpotsOpen int    `json:"advance_spots_open"`
	StaffSpotsOpen   int    `json:"staff_spots_open"`
	WalkinSpotsOpen  int    `json:"walkin_spots_open"`
	TotalSpotsOpen   int    `json:"total_spots_open"`
}

type Club struct {
	ClassName                string      `json:"class_name"`
	ID                       string      `json:"id"`
	Ident                    string      `json:"ident"`
	Name                     string      `json:"name"`
	CompanyIdent             string      `json:"company_ident"`
	Subline                  string      `json:"subline"`
	URL                      string      `json:"url"`
	PhoneNumber              string      `json:"phone_number"`
	AddressLine1             string      `json:"address_line1"`
	AddressLine2             string      `json:"address_line2"`
	AddressCity              string      `json:"address_city"`
	AddressState             string      `json:"address_state"`
	AddressZip               string      `json:"address_zip"`
	AddressCountry           string      `json:"address_country"`
	Lat                      float64     `json:"lat"`
	Lng                      float64     `json:"lng"`
	Timezone                 string      `json:"timezone"`
	StatusID                 int         `json:"status_id"`
	ShouldBeExcludedFromMaps bool        `json:"should_be_excluded_from_maps"`
	IsActive                 bool        `json:"is_active"`
	IsTempClosed             bool        `json:"is_temp_closed"`
	IsReservable             bool        `json:"is_reservable"`
	TotalMinutes             int         `json:"total_minutes"`
	WorkoutMinutes           int         `json:"workout_minutes"`
	CleaningMinutes          int         `json:"cleaning_minutes"`
	StaffSpotsPerTimeSlot    int         `json:"staff_spots_per_time_slot"`
	WalkinSpotsPerTimeSlot   int         `json:"walkin_spots_per_time_slot"`
	AdvanceSpotsPerTimeSlot  int         `json:"advance_spots_per_time_slot"`
	WalkinSpotsOpen          interface{} `json:"walkin_spots_open"`
	AdvanceMemberSpotsOpen   interface{} `json:"advance_member_spots_open"`
	AdvanceGuestSpotsOpen    interface{} `json:"advance_guest_spots_open"`
	Hours                    struct {
		Sunday []struct {
			Open                    string `json:"open"`
			Close                   string `json:"close"`
			Senior                  bool   `json:"senior"`
			WorkoutMinutes          int    `json:"workout_minutes"`
			CleaningMinutes         int    `json:"cleaning_minutes"`
			StaffSpotsPerTimeSlot   int    `json:"staff_spots_per_time_slot"`
			WalkinSpotsPerTimeSlot  int    `json:"walkin_spots_per_time_slot"`
			AdvanceSpotsPerTimeSlot int    `json:"advance_spots_per_time_slot"`
		} `json:"sunday"`
		Monday []struct {
			Open                    string `json:"open"`
			Close                   string `json:"close"`
			Senior                  bool   `json:"senior"`
			WorkoutMinutes          int    `json:"workout_minutes"`
			CleaningMinutes         int    `json:"cleaning_minutes"`
			StaffSpotsPerTimeSlot   int    `json:"staff_spots_per_time_slot"`
			WalkinSpotsPerTimeSlot  int    `json:"walkin_spots_per_time_slot"`
			AdvanceSpotsPerTimeSlot int    `json:"advance_spots_per_time_slot"`
		} `json:"monday"`
		Tuesday []struct {
			Open                    string `json:"open"`
			Close                   string `json:"close"`
			Senior                  bool   `json:"senior"`
			WorkoutMinutes          int    `json:"workout_minutes"`
			CleaningMinutes         int    `json:"cleaning_minutes"`
			StaffSpotsPerTimeSlot   int    `json:"staff_spots_per_time_slot"`
			WalkinSpotsPerTimeSlot  int    `json:"walkin_spots_per_time_slot"`
			AdvanceSpotsPerTimeSlot int    `json:"advance_spots_per_time_slot"`
		} `json:"tuesday"`
		Wednesday []struct {
			Open                    string `json:"open"`
			Close                   string `json:"close"`
			Senior                  bool   `json:"senior"`
			WorkoutMinutes          int    `json:"workout_minutes"`
			CleaningMinutes         int    `json:"cleaning_minutes"`
			StaffSpotsPerTimeSlot   int    `json:"staff_spots_per_time_slot"`
			WalkinSpotsPerTimeSlot  int    `json:"walkin_spots_per_time_slot"`
			AdvanceSpotsPerTimeSlot int    `json:"advance_spots_per_time_slot"`
		} `json:"wednesday"`
		Thursday []struct {
			Open                    string `json:"open"`
			Close                   string `json:"close"`
			Senior                  bool   `json:"senior"`
			WorkoutMinutes          int    `json:"workout_minutes"`
			CleaningMinutes         int    `json:"cleaning_minutes"`
			StaffSpotsPerTimeSlot   int    `json:"staff_spots_per_time_slot"`
			WalkinSpotsPerTimeSlot  int    `json:"walkin_spots_per_time_slot"`
			AdvanceSpotsPerTimeSlot int    `json:"advance_spots_per_time_slot"`
		} `json:"thursday"`
		Friday []struct {
			Open                    string `json:"open"`
			Close                   string `json:"close"`
			Senior                  bool   `json:"senior"`
			WorkoutMinutes          int    `json:"workout_minutes"`
			CleaningMinutes         int    `json:"cleaning_minutes"`
			StaffSpotsPerTimeSlot   int    `json:"staff_spots_per_time_slot"`
			WalkinSpotsPerTimeSlot  int    `json:"walkin_spots_per_time_slot"`
			AdvanceSpotsPerTimeSlot int    `json:"advance_spots_per_time_slot"`
		} `json:"friday"`
		Saturday []struct {
			Open                    string `json:"open"`
			Close                   string `json:"close"`
			Senior                  bool   `json:"senior"`
			WorkoutMinutes          int    `json:"workout_minutes"`
			CleaningMinutes         int    `json:"cleaning_minutes"`
			StaffSpotsPerTimeSlot   int    `json:"staff_spots_per_time_slot"`
			WalkinSpotsPerTimeSlot  int    `json:"walkin_spots_per_time_slot"`
			AdvanceSpotsPerTimeSlot int    `json:"advance_spots_per_time_slot"`
		} `json:"saturday"`
	} `json:"hours"`
	TimeSlots  []TimeSlot `json:"time_slots"`
	CreatedInt int64      `json:"created_int"`
	UpdatedInt int64      `json:"updated_int"`
	IsEditable bool       `json:"is_editable"`
}
