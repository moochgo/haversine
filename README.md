# haversine
formula: <br />
a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2) <br />
c = 2 ⋅ atan2( √a, √(1−a) ) <br />
d = R ⋅ c <br />
where	φ is latitude, λ is longitude, R is earth’s radius (mean radius = 6,371km)
