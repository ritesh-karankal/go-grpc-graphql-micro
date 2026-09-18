import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { ArrowRight, Sparkles, Truck, Shield, RefreshCw } from "lucide-react";
import { ProductCard } from "../components/ProductCard";
import { LoadingSpinner } from "../components/LoadingSpinner";
import { graphqlRequest, queries } from "../lib/graphql";
import type { Product } from "../types";

const features = [
  {
    icon: Truck,
    title: "Fast Delivery",
    description: "Orders processed instantly through our microservices backend.",
  },
  {
    icon: Shield,
    title: "Secure Checkout",
    description: "Account validation and product verification on every order.",
  },
  {
    icon: RefreshCw,
    title: "Real-time Catalog",
    description: "Product search powered by Elasticsearch for instant results.",
  },
];

export function Home() {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    graphqlRequest<{ products: Product[] }>(queries.products, {
      pagination: { skip: 0, take: 6 },
    })
      .then((data) => setProducts(data.products))
      .catch((err: Error) => setError(err.message))
      .finally(() => setLoading(false));
  }, []);

  return (
    <div>
      <section className="relative overflow-hidden bg-ink text-white">
        <div className="absolute inset-0 bg-[radial-gradient(circle_at_20%_50%,rgba(52,211,153,0.15),transparent_50%)]" />
        <div className="absolute inset-0 bg-[radial-gradient(circle_at_80%_20%,rgba(255,255,255,0.05),transparent_40%)]" />
        <div className="relative mx-auto max-w-7xl px-4 py-20 sm:px-6 sm:py-28 lg:px-8">
          <div className="max-w-2xl">
            <div className="inline-flex items-center gap-2 rounded-full border border-white/20 bg-white/10 px-4 py-1.5 text-sm backdrop-blur">
              <Sparkles className="h-4 w-4 text-brand-400" />
              Powered by GraphQL Microservices
            </div>
            <h1 className="mt-6 font-serif text-5xl leading-[1.1] tracking-tight sm:text-6xl lg:text-7xl">
              Curated goods for{" "}
              <span className="italic text-brand-400">modern</span> living
            </h1>
            <p className="mt-6 text-lg leading-relaxed text-slate-300">
              Discover thoughtfully selected products with seamless search, cart,
              and checkout — all connected to a Go gRPC microservices backend.
            </p>
            <div className="mt-8 flex flex-wrap gap-4">
              <Link
                to="/shop"
                className="inline-flex items-center gap-2 rounded-full bg-brand-400 px-6 py-3.5 text-sm font-semibold text-ink transition hover:bg-brand-300"
              >
                Shop Collection
                <ArrowRight className="h-4 w-4" />
              </Link>
              <Link
                to="/account"
                className="inline-flex items-center gap-2 rounded-full border border-white/30 px-6 py-3.5 text-sm font-medium transition hover:bg-white/10"
              >
                Create Account
              </Link>
            </div>
          </div>
        </div>
      </section>

      <section className="mx-auto max-w-7xl px-4 py-16 sm:px-6 lg:px-8">
        <div className="grid gap-6 sm:grid-cols-3">
          {features.map((feature) => (
            <div
              key={feature.title}
              className="rounded-2xl border border-slate-200 bg-white p-6 transition hover:shadow-lg hover:shadow-slate-100"
            >
              <div className="flex h-11 w-11 items-center justify-center rounded-xl bg-brand-50 text-brand-600">
                <feature.icon className="h-5 w-5" />
              </div>
              <h3 className="mt-4 font-semibold">{feature.title}</h3>
              <p className="mt-2 text-sm leading-relaxed text-muted">
                {feature.description}
              </p>
            </div>
          ))}
        </div>
      </section>

      <section className="mx-auto max-w-7xl px-4 pb-20 sm:px-6 lg:px-8">
        <div className="flex items-end justify-between gap-4">
          <div>
            <p className="text-sm font-medium uppercase tracking-wider text-brand-600">
              Featured
            </p>
            <h2 className="mt-2 font-serif text-3xl tracking-tight sm:text-4xl">
              Popular picks
            </h2>
          </div>
          <Link
            to="/shop"
            className="hidden sm:inline-flex items-center gap-1 text-sm font-medium text-muted transition hover:text-ink"
          >
            View all
            <ArrowRight className="h-4 w-4" />
          </Link>
        </div>

        {loading && <LoadingSpinner />}
        {error && (
          <div className="mt-8 rounded-2xl border border-red-200 bg-red-50 px-6 py-4 text-sm text-red-700">
            Could not load products: {error}. Make sure the backend is running at{" "}
            <code className="rounded bg-red-100 px-1">localhost:8000</code>.
          </div>
        )}
        {!loading && !error && products.length === 0 && (
          <div className="mt-8 rounded-2xl border border-dashed border-slate-300 bg-white px-6 py-12 text-center">
            <p className="text-muted">No products yet. Add some via GraphQL or the admin page.</p>
            <Link to="/shop" className="mt-4 inline-block text-sm font-medium text-brand-600">
              Go to shop
            </Link>
          </div>
        )}
        {!loading && products.length > 0 && (
          <div className="mt-8 grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
            {products.map((product, index) => (
              <ProductCard key={product.id} product={product} featured={index === 0} />
            ))}
          </div>
        )}
      </section>
    </div>
  );
}
