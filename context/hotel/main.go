package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := BookRoom(ctx); err != nil {
		fmt.Println(err)
	}
}

func BookRoom(ctx context.Context) error {
	select {
	case <-ctx.Done():
		fmt.Println("Booking cancelled. Time exceeded.")
		return ctx.Err()

	case <-time.After(2 * time.Second):
		fmt.Println("Room booked successfully.")
	}
	return nil
}
