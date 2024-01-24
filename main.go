package main

import "fmt"

func main() {
	m := map[string]int{}
	m["a"]--
	fmt.Println(m)
	m["b"]--
	fmt.Println(m)
}

//func main() {
//	eg, ctx := errgroup.WithContext(context.Background())
//
//	eg.Go(func() error {
//		time.Sleep(1000)
//		select {
//		case <-ctx.Done():
//			return nil
//		default:
//			fmt.Println("A")
//			return nil
//		}
//	})
//	eg.Go(func() error {
//		fmt.Println("B")
//		return errors.New("this is B")
//	})
//	eg.Go(func() error {
//		time.Sleep(1000)
//		select {
//		case <-ctx.Done():
//			return nil
//		default:
//			fmt.Println("C")
//			return nil
//		}
//	})
//	eg.Go(func() error {
//		time.Sleep(1000)
//		select {
//		case <-ctx.Done():
//			return nil
//		default:
//			fmt.Println("D")
//			return nil
//		}
//	})
//	eg.Go(func() error {
//		time.Sleep(1000)
//		select {
//		case <-ctx.Done():
//			return nil
//		default:
//			fmt.Println("E")
//			return nil
//		}
//	})
//	eg.Go(func() error {
//		fmt.Println("F")
//		return errors.New("this is F")
//	})
//
//	if err := eg.Wait(); err != nil {
//		fmt.Println(err)
//	}
//}
