package hole

import (
	"fmt"
	"strings"
)

func brainfuck() []Run {
	var tests []test = fixedBFCases()

	tests = append(tests, randomBFCase(5, 3, 8, -3, 2))
	tests = append(tests, randomBFCase(10, 5, 1, 1, -3))

	for i := 0; i < 3; i++ {
		jumpSize := randInt(5, 12)
		buckets := randInt(2, 8)
		initialBucketSize := randInt(1, 7)
		bucketSizeChange := randInt(-1, 2)
		charShift := randInt(-3, 3)
		tests = append(tests, randomBFCase(jumpSize, buckets, initialBucketSize, bucketSizeChange, charShift))
	}

	return outputTests(shuffle(tests))
}

func fixedBFCases() []test {
	return []test{
		test{"+++++++++++++++[>++>+++>++++>+++++>++++++>+++++++>++++++++<<<<<<<-]+++++++++++++++>>>++++++.>>+++++++.>>-----.<-.<<<<<<-----.",
			"Bash"},
		test{"+++++++++++++++[>++>+++>++++>+++++>++++++>+++++++>++++++++<<<<<<<-]+++++++++++++++>>>>-.>+++++++.>>--.<<.<+++++++++.>++.>>----.<.>--.++++.<<<<<<<-----.",
			"JavaScript"},
		test{"+++++++++++++++++++++++++[>++>+++>++++>+++++<<<<-]+++++++++++++++++++++++++>>+.>>--------.<---.<<<---------------.",
			"Lua"},
		test{"+++++++++++++++++++++++++[>++>+++>++++>+++++<<<<-]+++++++++++++++++++++++++>>+++++.>+.>-----------.------.<<<<---------------.",
			"Perl"},
		test{"+++++++++++++++[>++>+++>++++>+++++>++++++>+++++++>++++++++<<<<<<<-]+++++++++++++++>>>>+++++.>>----.>------.------.<<<<<<++.>>------.<<<-----.",
			"Perl 6"},
		test{"+++++++++++++++++++++[>++>+++>++++>+++++>++++++<<<<<-]+++++++++++++++++++++>>>----.--------.++++++++.<<<-----------.",
			"PHP"},
		test{"+++++++++++++++++++++[>++>+++>++++>+++++>++++++<<<<<-]+++++++++++++++++++++>>>----.>>-----.-----.<-.>-----.-.<<<<<-----------.",
			"Python"},
		test{"+++++++++++++++++++++[>++>+++>++++>+++++>++++++<<<<<-]+++++++++++++++++++++>>>--.>>---------.<-------.>++++.<<<<<-----------.",
			"Ruby"},
		test{"++++++++++++++++++[>++>+++>++++>+++++>++++++>+++++++<<<<<<-]++++++++++++++++++>>>-----.>>+++.<++++++++++.+.<<<----.>>++++.>>.---.<+.<<<<--------.",
			"Code Golf"},
		test{">>>>>>>>>>>>>>>>>>>>>>>>>>>>++++++++++++++++++++++++++[-<<[+<]+[>]>][<<[[-]-----<]>[>]>]<<[++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++<]>[.>]++++++++++.",
			"abcdefghijklmnopqrstuvwxyz"},
		test{"+++++[>+++++[>++>++>+++>+++>++++>++++<<<<<<-]<-]+++++[>>[>]<[+.<<]>[++.>>>]<[+.<]>[-.>>]<[-.<<<]>[.>]<[+.<]<-]++++++++++.",
			"eL34NfeOL454KdeJ44JOdefePK55gQ67ShfTL787KegJ77JTeghfUK88iV9:XjgYL:;:KfiJ::JYfijgZK;;k[<=]lh^L=>=KgkJ==J^gklh_K>>m`?@bnicL@A@KhmJ@@JchmnidKAA"},
		test{"++++++++++[>++++++++++>++++++++++++<<-]>--->++>+++++++++++++[<.-<.+>>-]++++[<<-------------------------.>>-]",
			"zaybxcwdveuftgshriqjpkolnmU<#"},
		test{"++++++++++[>+++++>+++++++>+<<<-]>+>++[-<.+>]>.",
			"3456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz"},
	}
}

func intToBFString(n int) string {
	if n < 1 {
		return strings.Repeat("-", -n)
	}
	return strings.Repeat("+", n)
}

func randomBFCase(jumpSize, buckets, initialBucketSize, bucketSizeChange, charShift int) test {
	// Only generates ASCII values in a range that was already used by fixed cases before random cases were added, in order to not break legacy solutions.
	const ASCIIMin = 50
	const ASCIIMax = 122

	bucketString := ""
	out := ""

	// Ensure all bucket sizes are strictly positive
	if bucketSizeChange*buckets+initialBucketSize < 1 {
		bucketSizeChange = 1
	}

	// Populate buckets
	bucketSize := initialBucketSize
	for i := 0; i < buckets; i++ {
		// Choose a base for this bucket that ensures all chars remain in the printable ASCII range
		totalShift := charShift * (bucketSize - 1)
		minChar := ASCIIMin
		maxChar := ASCIIMax
		if totalShift > 0 {
			maxChar -= totalShift
		} else {
			minChar -= totalShift
		}
		base := randInt(minChar/jumpSize+1, maxChar/jumpSize)
		bucketString += ">" + intToBFString(base)

		// Add all the chars in the bucket to the solution
		char := base * jumpSize
		for j := 0; j < bucketSize; j++ {
			out += fmt.Sprintf("%c", char)
			char += charShift
		}

		bucketSize += bucketSizeChange
	}

	in := fmt.Sprintf(">%s[%s[<]>-]<%s>>[<<[>+>.%s<<-]>%s>[-]>[<]<<[>++++++++++.>]>>>]",
		intToBFString(jumpSize),
		bucketString,
		intToBFString(initialBucketSize),
		intToBFString(charShift),
		intToBFString(bucketSizeChange))

	return test{in, out}
}
