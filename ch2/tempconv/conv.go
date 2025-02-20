package tempconv

// CToF преобразует температуру по Цельсию в температуру по Фаренгейту
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.15) }

func MillToCm(mill Millimetr) Centimetr { return Centimetr(mill / 10) }
func CmToM(cm Centimetr) Metr           { return Metr(cm / 100) }
func MillToM(mill Millimetr) Metr       { return Metr(mill / 1000) }

func GrToKg(gr Gram) Kilogram { return Kilogram(gr / 1000) }
func KgToT(kg Kilogram) Ton   { return Ton(kg / 1000) }
func GrToT(gr Gram) Ton       { return Ton(gr / 1000000) }
