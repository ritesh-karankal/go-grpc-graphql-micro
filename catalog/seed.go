package catalog

import (
	"context"
	"log"
)

func SeedSampleProducts(ctx context.Context, s Service) error {
	products, err := s.GetProducts(ctx, 0, 100)
	if err != nil {
		return err
	}

	existing := make(map[string]struct{}, len(products))
	for _, product := range products {
		existing[product.Name] = struct{}{}
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
		{
			name:        "Wireless Mouse",
			description: "Ergonomic wireless mouse with adjustable sensitivity, silent clicks, and 18-month battery life.",
			price:       39.99,
		},
		{
			name:        "4K Monitor",
			description: "27-inch 4K UHD monitor with an IPS panel, HDR support, and USB-C connectivity.",
			price:       329.99,
		},
		{
			name:        "Portable SSD",
			description: "1TB portable solid-state drive with USB-C connectivity and transfer speeds up to 1,050 MB/s.",
			price:       99.99,
		},
		{
			name:        "USB-C GaN Charger",
			description: "65W compact wall charger with two USB-C ports and fast charging for laptops and mobile devices.",
			price:       44.99,
		},
		{
			name:        "Bluetooth Speaker",
			description: "Portable waterproof speaker with 360-degree sound and up to twelve hours of battery life.",
			price:       69.99,
		},
		{
			name:        "USB Microphone",
			description: "Plug-and-play condenser microphone with a cardioid pickup pattern and headphone monitoring.",
			price:       84.99,
		},
		{
			name:        "Gaming Controller",
			description: "Wireless controller with textured grips, responsive triggers, and compatibility with PC and console.",
			price:       54.99,
		},
	}

	for _, sample := range samples {
		if _, ok := existing[sample.name]; ok {
			continue
		}
		if _, err := s.PostProduct(ctx, sample.name, sample.description, sample.price); err != nil {
			return err
		}
	}

	log.Println("Seeded sample products")
	return nil
}
