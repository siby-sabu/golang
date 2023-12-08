module multimodulesample/module2

go 1.21.3

replace multimodulesample/module1 => ../module1

require multimodulesample/module1 v0.0.0-00010101000000-000000000000
