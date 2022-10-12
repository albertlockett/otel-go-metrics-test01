example to show difference between golang and js histogram

```
histogram.record(95)
histogram.record(95)
histogram.record(95)
histogram.record(245)
histogram.record(245)
histogram.record(495)
```

here is the golang output:
```
curl -XGET http://localhost:2222/metrics
# HELP my_histogram 
# TYPE my_histogram histogram
my_histogram_bucket{le="0"} 0
my_histogram_bucket{le="5"} 0
my_histogram_bucket{le="10"} 0
my_histogram_bucket{le="25"} 0
my_histogram_bucket{le="50"} 0
my_histogram_bucket{le="75"} 0
my_histogram_bucket{le="100"} 3
my_histogram_bucket{le="250"} 2
my_histogram_bucket{le="500"} 1
my_histogram_bucket{le="750"} 0
my_histogram_bucket{le="1000"} 0
my_histogram_bucket{le="2500"} 0
my_histogram_bucket{le="5000"} 0
my_histogram_bucket{le="7500"} 0
my_histogram_bucket{le="10000"} 0
my_histogram_bucket{le="+Inf"} 6
my_histogram_sum 1270
my_histogram_count 6
```

and here is the JS output:
```
curl -XGET localhost:9464/metrics
# HELP my_histogram description missing
# TYPE my_histogram histogram
my_histogram_count 6 1665572463653
my_histogram_sum 1270 1665572463653
my_histogram_bucket{le="0"} 0 1665572463653
my_histogram_bucket{le="5"} 0 1665572463653
my_histogram_bucket{le="10"} 0 1665572463653
my_histogram_bucket{le="25"} 0 1665572463653
my_histogram_bucket{le="50"} 0 1665572463653
my_histogram_bucket{le="75"} 0 1665572463653
my_histogram_bucket{le="100"} 3 1665572463653
my_histogram_bucket{le="250"} 5 1665572463653
my_histogram_bucket{le="500"} 6 1665572463653
my_histogram_bucket{le="1000"} 6 1665572463653
my_histogram_bucket{le="+Inf"} 6 1665572463653
```
