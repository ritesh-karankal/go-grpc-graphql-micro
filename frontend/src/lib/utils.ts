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

export function formatDate(date: string) {
  return new Intl.DateTimeFormat("en-US", {
    dateStyle: "medium",
    timeStyle: "short",
  }).format(new Date(date));
}
