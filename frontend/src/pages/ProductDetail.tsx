import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { ArrowLeft, Minus, Plus, ShoppingBag, Check } from "lucide-react";
import { LoadingSpinner } from "../components/LoadingSpinner";
import { graphqlRequest, queries } from "../lib/graphql";
import { useCart, formatPrice } from "../lib/cart";
import { productImage } from "../lib/utils";
import type { Product } from "../types";

export function ProductDetail() {
  const { id } = useParams<{ id: string }>();
  const { addItem } = useCart();
  const [product, setProduct] = useState<Product | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [quantity, setQuantity] = useState(1);
  const [added, setAdded] = useState(false);

  useEffect(() => {
    if (!id) return;

    graphqlRequest<{ products: Product[] }>(queries.products, { id })
      .then((data) => setProduct(data.products[0] ?? null))
      .catch((err: Error) => setError(err.message))
      .finally(() => setLoading(false));
  }, [id]);

  function handleAddToCart() {
    if (!product) return;
    addItem(product, quantity);
    setAdded(true);
    setTimeout(() => setAdded(false), 2000);
  }

  if (loading) return <LoadingSpinner />;

  if (error || !product) {
    return (
      <div className="mx-auto max-w-7xl px-4 py-20 text-center sm:px-6">
        <h1 className="text-2xl font-semibold">Product not found</h1>
        <p className="mt-2 text-muted">{error ?? "This product may have been removed."}</p>
        <Link to="/shop" className="mt-6 inline-block text-sm font-medium text-brand-600">
          Back to shop
        </Link>
      </div>
    );
  }

  const gradient = productImage(product.name);

  return (
    <div className="mx-auto max-w-7xl px-4 py-10 sm:px-6 lg:px-8">
      <Link
        to="/shop"
        className="inline-flex items-center gap-2 text-sm text-muted transition hover:text-ink"
      >
        <ArrowLeft className="h-4 w-4" />
        Back to shop
      </Link>

      <div className="mt-8 grid gap-10 lg:grid-cols-2 lg:gap-16">
        <div
          className={`relative flex aspect-square items-center justify-center overflow-hidden rounded-3xl bg-gradient-to-br ${gradient}`}
        >
          <div className="absolute inset-0 bg-[radial-gradient(circle_at_30%_20%,rgba(255,255,255,0.8),transparent_50%)]" />
          <span className="relative font-serif text-[12rem] leading-none text-ink/15">
            {product.name.charAt(0).toUpperCase()}
          </span>
        </div>

        <div className="flex flex-col justify-center">
          <p className="text-sm font-medium uppercase tracking-wider text-brand-600">
            Product
          </p>
          <h1 className="mt-2 font-serif text-4xl tracking-tight sm:text-5xl">
            {product.name}
          </h1>
          <p className="mt-6 text-lg leading-relaxed text-muted">{product.description}</p>
          <p className="mt-8 text-3xl font-semibold">{formatPrice(product.price)}</p>

          <div className="mt-8 flex items-center gap-4">
            <div className="flex items-center rounded-full border border-slate-200">
              <button
                type="button"
                onClick={() => setQuantity((q) => Math.max(1, q - 1))}
                className="flex h-11 w-11 items-center justify-center text-muted transition hover:text-ink"
                aria-label="Decrease quantity"
              >
                <Minus className="h-4 w-4" />
              </button>
              <span className="w-10 text-center font-medium">{quantity}</span>
              <button
                type="button"
                onClick={() => setQuantity((q) => q + 1)}
                className="flex h-11 w-11 items-center justify-center text-muted transition hover:text-ink"
                aria-label="Increase quantity"
              >
                <Plus className="h-4 w-4" />
              </button>
            </div>

            <button
              type="button"
              onClick={handleAddToCart}
              className="inline-flex flex-1 items-center justify-center gap-2 rounded-full bg-ink px-6 py-3 text-sm font-semibold text-white transition hover:bg-slate-800 active:scale-[0.98]"
            >
              {added ? (
                <>
                  <Check className="h-4 w-4" />
                  Added to cart
                </>
              ) : (
                <>
                  <ShoppingBag className="h-4 w-4" />
                  Add to cart
                </>
              )}
            </button>
          </div>

          <Link
            to="/cart"
            className="mt-4 text-center text-sm font-medium text-brand-600 hover:underline"
          >
            View cart
          </Link>
        </div>
      </div>
    </div>
  );
}
