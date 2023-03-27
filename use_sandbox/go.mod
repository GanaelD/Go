module example.com/use_sandbox

go 1.20

require example.com/challenges v0.0.0-00010101000000-000000000000

replace example.com/sandbox => ../sandbox

replace example.com/interfaces => ../interfaces

replace example.com/challenges => ../challenges
