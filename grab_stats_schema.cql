// {"DTO":"$GOPATH/src/github.com/herlegs/GoTest/playground/playground.go.PaxCompletedBookingDTO","GenerateTime":"2020-06-11T16:59:05.397525+08:00"}
CREATE TABLE IF NOT EXISTS grab_stats.passenger_completed_bookings (
    passenger_id bigint,
    completed_time timestamp,
    vehicle_type_id bigint,
    booking_code text,
    PRIMARY KEY (passenger_id, completed_time)
) WITH CLUSTERING ORDER BY (completed_time desc);