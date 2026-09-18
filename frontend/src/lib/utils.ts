export function cn(...classes: Array<string | false | null | undefined>) {
  return classes.filter(Boolean).join(" ");
}

export function productImage(name: string) {
  const hash = name.split("").reduce((acc, char) => acc + char.charCodeAt(0), 0);
  const gradients = [
    "from-emerald-100 via-teal-50 to-cyan-100",
    "from-violet-100 via-purple-50 to-fuchsia-100",
    "from-amber-100 via-orange-50 to-rose-100",
    "from-sky-100 via-blue-50 to-indigo-100",
    "from-lime-100 via-green-50 to-emerald-100",
    "from-rose-100 via-pink-50 to-purple-100",
  ];
  return gradients[hash % gradients.length];
}

// Return a local curated photo path for a product. Uses a small bundled set of studio images.
export function productPhoto(name: string, size = 800) {
  // Local images placed in frontend/public/assets/images/photo-{1..6}.jpg
  const images = [
    "/assets/images/photo-1.jpg",
    "/assets/images/photo-2.jpg",
    "/assets/images/photo-3.jpg",
    "/assets/images/photo-4.jpg",
    "/assets/images/photo-5.jpg",
    "/assets/images/photo-6.jpg",
  ];

  // simple stable hash to pick an image per product name
  const hash = name.split("").reduce((acc, char) => acc + char.charCodeAt(0), 0);
  return images[hash % images.length];
}

export function formatDate(date: string) {
  return new Intl.DateTimeFormat("en-US", {
    dateStyle: "medium",
    timeStyle: "short",
  }).format(new Date(date));
}
