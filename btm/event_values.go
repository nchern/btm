package main

func (e *event) toFields() map[string]interface{} {
	return map[string]interface{}{
		"date":                   e.raw[0],
		"device":                 e.raw[1],
		"direction.":             e.raw[2],
		"hex":                    e.raw[3],
		"iridium_lat;":           e.raw[4],
		"iridium_long;":          e.raw[5],
		"month":                  e.raw[6],
		"day":                    e.raw[7],
		"hour":                   e.raw[8],
		"minute":                 e.raw[9],
		"second":                 e.raw[10],
		"latitude":               e.raw[11],
		"longitude":              e.raw[12],
		"heading":                e.raw[13],
		"pitch":                  e.raw[14],
		"roll":                   e.raw[15],
		"thruster_v":             e.raw[16],
		"thruster_a":             e.raw[17],
		"hotel_v":                e.raw[18],
		"hotel_a":                e.raw[19],
		"solar_v":                e.raw[20],
		"solar_a":                e.raw[21],
		"throttle":               e.raw[22],
		"rpm":                    e.raw[23],
		"rudder_angle":           e.raw[24],
		"next_waypoint_number":   e.raw[25],
		"next_waypoint_lat":      e.raw[26],
		"next_waypoint_long":     e.raw[27],
		"target_heading":         e.raw[28],
		"cross_track_error":      e.raw[29],
		"next_waypoint_distance": e.raw[30],
		"average_speed":          e.raw[31],
		"uptime_mins":            e.raw[32],
	}
}
