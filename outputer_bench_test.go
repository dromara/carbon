package carbon

import (
	"sync"
	"testing"
)

func BenchmarkCarbon_GoString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.GoString()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.GoString()
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.GoString()
			}
		})
	})
}

func BenchmarkCarbon_ToString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToMonthString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToMonthString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToMonthString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToMonthString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortMonthString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortMonthString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortMonthString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortMonthString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToWeekString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToWeekString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToWeekString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToWeekString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortWeekString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortWeekString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortWeekString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortWeekString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToDayDateTimeString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToDayDateTimeString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToDayDateTimeString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToDayDateTimeString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToDateTimeString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToDateTimeString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToDateTimeString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToDateTimeString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToDateTimeMilliString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToDateTimeMilliString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToDateTimeMilliString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToDateTimeMilliString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToDateTimeMicroString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToDateTimeMicroString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToDateTimeMicroString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToDateTimeMicroString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToDateTimeNanoString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToDateTimeNanoString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToDateTimeNanoString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToDateTimeNanoString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortDateTimeString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortDateTimeString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortDateTimeString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortDateTimeString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortDateTimeMilliString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortDateTimeMilliString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortDateTimeMilliString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortDateTimeMilliString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortDateTimeMicroString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortDateTimeMicroString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortDateTimeMicroString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortDateTimeMicroString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortDateTimeNanoString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortDateTimeNanoString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortDateTimeNanoString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortDateTimeNanoString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToDateString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToDateString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToDateString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToDateString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToDateMilliString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToDateMilliString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToDateMilliString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToDateMilliString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToDateMicroString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToDateMicroString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToDateMicroString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToDateMicroString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToDateNanoString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToDateNanoString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToDateNanoString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToDateNanoString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortDateString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortDateString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortDateString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortDateString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortDateMilliString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortDateMilliString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortDateMilliString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortDateMilliString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortDateMicroString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortDateMicroString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortDateMicroString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortDateMicroString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortDateNanoString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortDateNanoString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortDateNanoString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortDateNanoString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToTimeString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToTimeString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToTimeString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToTimeString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToTimeMilliString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToTimeMilliString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToTimeMilliString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToTimeMilliString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToTimeMicroString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToTimeMicroString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToTimeMicroString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToTimeMicroString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToTimeNanoString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToTimeNanoString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToTimeNanoString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToTimeNanoString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortTimeString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortTimeString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortTimeString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortTimeString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortTimeMilliString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortTimeMilliString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortTimeMilliString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortTimeMilliString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortTimeMicroString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortTimeMicroString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortTimeMicroString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortTimeMicroString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToShortTimeNanoString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToShortTimeNanoString(PRC)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToShortTimeNanoString(PRC)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToShortTimeNanoString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToAtomString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToAtomString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToAtomString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToAtomString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToAnsicString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToAnsicString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToAnsicString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToAnsicString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToCookieString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToCookieString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToCookieString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToCookieString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRssString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRssString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRssString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRssString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToW3cString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToW3cString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToW3cString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToW3cString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToUnixDateString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToUnixDateString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToUnixDateString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToUnixDateString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRubyDateString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRubyDateString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRubyDateString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRubyDateString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToKitchenString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToKitchenString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToKitchenString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToKitchenString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToHttpString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToHttpString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToHttpString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToHttpString(PRC)
			}
		})
	})
}
func BenchmarkCarbon_ToIso8601String(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToIso8601String(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToIso8601String(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToIso8601String(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToIso8601MilliString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToIso8601MilliString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToIso8601MilliString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToIso8601MilliString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToIso8601NanoString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToIso8601NanoString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToIso8601NanoString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToIso8601NanoString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToIso8601ZuluString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToIso8601ZuluString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToIso8601ZuluString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToIso8601ZuluString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToIso8601ZuluMilliString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToIso8601ZuluMilliString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToIso8601ZuluMilliString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToIso8601ZuluMilliString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToIso8601ZuluMicroString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToIso8601ZuluMicroString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToIso8601ZuluMicroString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToIso8601ZuluMicroString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToIso8601ZuluNanoString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToIso8601ZuluNanoString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToIso8601ZuluNanoString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToIso8601ZuluNanoString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc822String(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc822String(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc822String(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc822String(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc822zString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc822zString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc822zString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc822zString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc850String(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc850String(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc850String(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc850String(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc1036String(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc1036String(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc1036String(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc1036String(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc1123String(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc1123String(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc1123String(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc1123String(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc1123zString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc1123zString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc1123zString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc1123zString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc2822String(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc2822String(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc2822String(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc2822String(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc3339String(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc3339String(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc3339String(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc3339String(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc3339MilliString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc3339MilliString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc3339MilliString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc3339MilliString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc3339MicroString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc3339MicroString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc3339MicroString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc3339MicroString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc3339NanoString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc3339NanoString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc3339NanoString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc3339NanoString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToRfc7231String(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToRfc7231String(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToRfc7231String(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToRfc7231String(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToFormattedDateString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToFormattedDateString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToFormattedDateString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToFormattedDateString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_ToFormattedDayDateString(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.ToFormattedDayDateString(PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.ToFormattedDayDateString(PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.ToFormattedDayDateString(PRC)
			}
		})
	})
}

func BenchmarkCarbon_Layout(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.Layout(DateTimeLayout, PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.Layout(DateTimeLayout, PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.Layout(DateTimeLayout, PRC)
			}
		})
	})
}

func BenchmarkCarbon_Format(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			c.Format(DateTimeLayout, PRC)
		}
	})
	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N/10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.Format(DateTimeLayout, PRC)
			}()
		}
		wg.Wait()
	})
	b.Run("parallel", func(b *testing.B) {
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.Format(DateTimeLayout, PRC)
			}
		})
	})
}
