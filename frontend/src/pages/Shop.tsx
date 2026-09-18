import { useEffect, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { Package, Search } from "lucide-react";
import { ProductCard } from "../components/ProductCard";
import { LoadingSpinner } from "../components/LoadingSpinner";
import { EmptyState } from "../components/EmptyState";
import { graphqlRequest, queries } from "../lib/graphql";
import type { Product } from "../types";

export function Shop() {
  const [searchParams, setSearchParams] = useSearchParams();
  const queryParam = searchParams.get("q") ?? "";
  const [search, setSearch] = useState(queryParam);
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setSearch(queryParam);
  }, [queryParam]);

  useEffect(() => {
    setLoading(true);
    setError(null);

    const variables: Record<string, unknown> = {
      pagination: { skip: 0, take: 50 },
    };
    if (queryParam) {
      variables.query = queryParam;
    }

    graphqlRequest<{ products: Product[] }>(queries.products, variables)
      .then((data) => setProducts(data.products))
      .catch((err: Error) => setError(err.message))
      .finally(() => setLoading(false));
  }, [queryParam]);

  function handleSearch(event: React.FormEvent) {
    event.preventDefault();
    const trimmed = search.trim();
    if (trimmed) {
      setSearchParams({ q: trimmed });
    } else {
      setSearchParams({});
    }
  }

  return (
    <div className="mx-auto max-w-7xl px-4 py-10 sm:px-6 lg:px-8">
      <div className="max-w-2xl">
        <p className="text-sm font-medium uppercase tracking-wider text-brand-600">
          Catalog
        </p>
        <h1 className="mt-2 font-serif text-4xl tracking-tight sm:text-5xl">
          Shop all products
        </h1>
        <p className="mt-4 text-muted">
          Search the catalog powered by Elasticsearch. Results update in real time.
        </p>
      </div>

      <form onSubmit={handleSearch} className="mt-8 max-w-xl">
        <div className="relative">
          <Search className="absolute left-4 top-1/2 h-5 w-5 -translate-y-1/2 text-muted" />
          <input
            type="search"
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            placeholder="Search by name or description..."
            className="w-full rounded-2xl border border-slate-200 bg-white py-3.5 pl-12 pr-4 text-sm shadow-sm outline-none transition focus:border-brand-400 focus:ring-2 focus:ring-brand-100"
          />
        </div>
      </form>

      {queryParam && (
        <p className="mt-4 text-sm text-muted">
          Showing results for &ldquo;{queryParam}&rdquo; &middot;{" "}
          <button
            type="button"
            onClick={() => setSearchParams({})}
            className="font-medium text-brand-600 hover:underline"
          >
            Clear search
          </button>
        </p>
      )}

      {loading && <LoadingSpinner />}
      {error && (
        <div className="mt-8 rounded-2xl border border-red-200 bg-red-50 px-6 py-4 text-sm text-red-700">
          {error}
        </div>
      )}
      {!loading && !error && products.length === 0 && (
        <div className="mt-8">
          <EmptyState
            icon={Package}
            title="No products found"
            description={
              queryParam
                ? "Try a different search term or browse the full catalog."
                : "The catalog is empty. Create products via GraphQL to get started."
            }
            actionLabel="Clear search"
            actionTo="/shop"
          />
        </div>
      )}
      {!loading && products.length > 0 && (
        <div className="mt-10 grid gap-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
          {products.map((product) => (
            <ProductCard key={product.id} product={product} />
          ))}
        </div>
      )}
    </div>
  );
}
