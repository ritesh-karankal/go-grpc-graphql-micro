import { Link } from "react-router-dom";
import { Minus, Plus, ShoppingBag, Trash2 } from "lucide-react";
import { useCart, formatPrice } from "../lib/cart";
import { EmptyState } from "../components/EmptyState";
import { productImage } from "../lib/utils";

export function Cart() {
  const { items, subtotal, updateQuantity, removeItem } = useCart();

  if (items.length === 0) {
    return (
      <div className="mx-auto max-w-7xl px-4 py-10 sm:px-6 lg:px-8">
        <h1 className="font-serif text-4xl tracking-tight">Your cart</h1>
        <div className="mt-8">
          <EmptyState
            icon={ShoppingBag}
            title="Your cart is empty"
            description="Browse our collection and add products you love."
            actionLabel="Start shopping"
            actionTo="/shop"
          />
        </div>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-7xl px-4 py-10 sm:px-6 lg:px-8">
      <h1 className="font-serif text-4xl tracking-tight">Your cart</h1>
      <p className="mt-2 text-muted">{items.length} item{items.length !== 1 ? "s" : ""}</p>

      <div className="mt-10 grid gap-10 lg:grid-cols-3">
        <div className="lg:col-span-2 space-y-4">
          {items.map(({ product, quantity }) => (
            <div
              key={product.id}
              className="flex gap-4 rounded-2xl border border-slate-200 bg-white p-4 sm:gap-6 sm:p-5"
            >
              <div
                className={`flex h-24 w-24 shrink-0 items-center justify-center rounded-xl bg-gradient-to-br ${productImage(product.name)}`}
              >
                <span className="font-serif text-3xl text-ink/20">
                  {product.name.charAt(0).toUpperCase()}
                </span>
              </div>

              <div className="flex flex-1 flex-col">
                <div className="flex justify-between gap-4">
                  <div>
                    <Link
                      to={`/product/${product.id}`}
                      className="font-semibold transition hover:text-brand-600"
                    >
                      {product.name}
                    </Link>
                    <p className="mt-1 line-clamp-1 text-sm text-muted">
                      {product.description}
                    </p>
                  </div>
                  <p className="shrink-0 font-semibold">
                    {formatPrice(product.price * quantity)}
                  </p>
                </div>

                <div className="mt-auto flex items-center justify-between pt-4">
                  <div className="flex items-center rounded-full border border-slate-200">
                    <button
                      type="button"
                      onClick={() => updateQuantity(product.id, quantity - 1)}
                      className="flex h-9 w-9 items-center justify-center text-muted hover:text-ink"
                      aria-label="Decrease quantity"
                    >
                      <Minus className="h-3.5 w-3.5" />
                    </button>
                    <span className="w-8 text-center text-sm font-medium">{quantity}</span>
                    <button
                      type="button"
                      onClick={() => updateQuantity(product.id, quantity + 1)}
                      className="flex h-9 w-9 items-center justify-center text-muted hover:text-ink"
                      aria-label="Increase quantity"
                    >
                      <Plus className="h-3.5 w-3.5" />
                    </button>
                  </div>

                  <button
                    type="button"
                    onClick={() => removeItem(product.id)}
                    className="flex items-center gap-1.5 text-sm text-muted transition hover:text-red-600"
                  >
                    <Trash2 className="h-4 w-4" />
                    Remove
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>

        <div className="h-fit rounded-2xl border border-slate-200 bg-white p-6">
          <h2 className="text-lg font-semibold">Order summary</h2>
          <div className="mt-6 space-y-3 text-sm">
            <div className="flex justify-between">
              <span className="text-muted">Subtotal</span>
              <span className="font-medium">{formatPrice(subtotal)}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-muted">Shipping</span>
              <span className="font-medium text-brand-600">Free</span>
            </div>
            <div className="border-t border-slate-200 pt-3 flex justify-between text-base">
              <span className="font-semibold">Total</span>
              <span className="font-semibold">{formatPrice(subtotal)}</span>
            </div>
          </div>

          <Link
            to="/checkout"
            className="mt-6 flex w-full items-center justify-center rounded-full bg-ink py-3.5 text-sm font-semibold text-white transition hover:bg-slate-800"
          >
            Proceed to checkout
          </Link>

          <Link
            to="/shop"
            className="mt-3 block text-center text-sm text-muted transition hover:text-ink"
          >
            Continue shopping
          </Link>
        </div>
      </div>
    </div>
  );
}
