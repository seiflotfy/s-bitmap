# s-bitmap
S-Bitmap: Distinct Counting with a Self-Learning Bitmap

The S-bitmap is a bitmap obtained via a novel adaptive sampling process, where the bits corresponding to the sampled items are set to 1, and the sampling rates are learned from the number of distinct items already passed and reduced sequentially as more bits are set to 1. A unique property of S-bitmap is that its relative estimation error is truly stabilized, i.e. invariant to unknown cardinalities in a prescribed range. This paper demonstrates through both theoretical and empirical studies that with a given memory requirement, S-bitmap is not only uniformly reliable but more accurate than state-of-the-art algorithms such as the multiresolution bitmap and HyperLogLog algorithms (not HyperLogLog++) under common practice settings.

For details about the algorithm and citations please use this article for now:
["Distinct Counting with a Self-Learning Bitmap" by Aiyou Chen and Jin Cao](http://ect.bell-labs.com/who/aychen/sbitmap4p.pdf)

## Note:
A portion of this code has been translated from the following [implementation](https://github.com/travisbrady/self-learning-bitmap)

##Example usage:
```go

import "github.com/seiflotfy/s-bitmap"

sb := sbitmap.NewDefault()

// Add 1000 different strings
for i:=0; i<1000; i++ {
	sb.Update([]byte{"foobar" + strconv.Itoa(i)})
}


count := sb.Estimate()
// count ~= 1000 (+/- 0.8%)
```

##Todo:
- [ ] Add Merge functionality