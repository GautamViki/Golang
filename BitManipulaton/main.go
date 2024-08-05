package main

import "bitsManipulation/handler"

func main() {
	bits := handler.NewBitManipulationHandler()
	bits.ConvertDecimalToBinary()
	bits.ConvertBinaryToDecimal()
	bits.SwapTwoNumber()
	bits.CheckNthBitSet()
	bits.ConvertNthBitToSet()
}
