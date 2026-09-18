package catalog

import (
	"context"
	"log"
)

func SeedSampleProducts(ctx context.Context, s Service) error {
	products, err := s.GetProducts(ctx, 0, 1)
	if err != nil {
		return err
	}
	if len(products) > 0 {
		return nil
	}

	samples := []struct {
		name        string
		description string
		price       float64
	}{
		{
			name:        "Wireless Headphones",
			description: "Premium noise-cancelling over-ear headphones with 30-hour battery life.",
			price:       149.99,
		},
		{
			name:        "Mechanical Keyboard",
			description: "Compact 75% layout keyboard with hot-swappable switches and RGB backlight.",
			price:       89.99,
		},
		{
			name:        "USB-C Hub",
			description: "7-in-1 adapter with HDMI, SD card reader, and 100W power delivery.",
			price:       49.99,
		},
		{
			name:        "Monitor Stand",
			description: "Aluminum riser with cable management and extra storage shelf.",
			price:       79.99,
		},
		{
			name:        "Laptop Sleeve",
			description: "Water-resistant neoprene sleeve with soft microfiber lining.",
			price:       34.99,
		},
		{
			name:        "HD Webcam",
			description: "1080p webcam with auto-focus, built-in microphone, and privacy cover.",
			price:       59.99,
		},
	}

	for _, sample := range samples {
		if _, err := s.PostProduct(ctx, sample.name, sample.description, sample.price); err != nil {
			return err
		}
	}

	log.Println("Seeded sample products")
	return nil
}
