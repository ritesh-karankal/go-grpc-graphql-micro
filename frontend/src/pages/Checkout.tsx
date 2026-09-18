import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { CheckCircle2, Loader2 } from "lucide-react";
import { useCart, formatPrice } from "../lib/cart";
import { graphqlRequest, queries } from "../lib/graphql";
import type { Order } from "../types";

export function Checkout() {
  const navigate = useNavigate();
  const { items, subtotal, accountId, accountName, clearCart } = useCart();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [order, setOrder] = useState<Order | null>(null);

  if (items.length === 0 && !order) {
    return (
      <div className="mx-auto max-w-lg px-4 py-20 text-center sm:px-6">
        <h1 className="text-2xl font-semibold">Nothing to checkout</h1>
        <p className="mt-2 text-muted">Add items to your cart first.</p>
        <Link to="/shop" className="mt-6 inline-block text-sm font-medium text-brand-600">
          Go to shop
        </Link>
      </div>
    );
  }

  if (order) {
    return (
      <div className="mx-auto max-w-lg px-4 py-20 text-center sm:px-6">
        <div className="mx-auto flex h-16 w-16 items-center justify-center rounded-full bg-brand-50">
          <CheckCircle2 className="h-8 w-8 text-brand-600" />
        </div>
        <h1 className="mt-6 text-2xl font-semibold">Order confirmed!</h1>
        <p className="mt-2 text-muted">
          Thank you{accountName ? `, ${accountName}` : ""}. Your order has been placed.
        </p>
        <div className="mt-8 rounded-2xl border border-slate-200 bg-white p-6 text-left">
          <p className="text-sm text-muted">Order ID</p>
          <p className="mt-1 font-mono text-sm">{order.id}</p>
          <p className="mt-4 text-sm text-muted">Total</p>
          <p className="mt-1 text-2xl font-semibold">{formatPrice(order.totalPrice)}</p>
        </div>
        <div className="mt-8 flex flex-col gap-3 sm:flex-row sm:justify-center">
          <Link
            to="/orders"
            className="rounded-full bg-ink px-6 py-3 text-sm font-semibold text-white transition hover:bg-slate-800"
          >
            View orders
          </Link>
          <Link
            to="/shop"
            className="rounded-full border border-slate-200 px-6 py-3 text-sm font-medium transition hover:bg-slate-50"
          >
            Continue shopping
          </Link>
        </div>
      </div>
    );
  }

  async function handleCheckout() {
    if (!accountId) {
      setError("Please sign in or create an account before checkout.");
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const data = await graphqlRequest<{ createOrder: Order }>(queries.createOrder, {
        order: {
          accountId,
          products: items.map((item) => ({
            id: item.product.id,
            quantity: item.quantity,
          })),
        },
      });

      setOrder(data.createOrder);
      clearCart();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Checkout failed");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="mx-auto max-w-7xl px-4 py-10 sm:px-6 lg:px-8">
      <h1 className="font-serif text-4xl tracking-tight">Checkout</h1>

      <div className="mt-10 grid gap-10 lg:grid-cols-2">
        <div>
          <h2 className="text-lg font-semibold">Account</h2>
          {accountId ? (
            <div className="mt-4 rounded-2xl border border-brand-200 bg-brand-50 p-5">
              <p className="text-sm text-brand-700">Signed in as</p>
              <p className="mt-1 font-semibold">{accountName}</p>
            </div>
          ) : (
            <div className="mt-4 rounded-2xl border border-amber-200 bg-amber-50 p-5">
              <p className="text-sm text-amber-800">
                You need an account to place an order.
              </p>
              <Link
                to="/account"
                className="mt-3 inline-block text-sm font-semibold text-amber-900 underline"
              >
                Sign in or create account
              </Link>
            </div>
          )}

          <h2 className="mt-8 text-lg font-semibold">Items</h2>
          <ul className="mt-4 space-y-3">
            {items.map(({ product, quantity }) => (
              <li
                key={product.id}
                className="flex justify-between rounded-xl border border-slate-200 bg-white px-4 py-3 text-sm"
              >
                <span>
                  {product.name} &times; {quantity}
                </span>
                <span className="font-medium">
                  {formatPrice(product.price * quantity)}
                </span>
              </li>
            ))}
          </ul>
        </div>

        <div className="h-fit rounded-2xl border border-slate-200 bg-white p-6">
          <h2 className="text-lg font-semibold">Payment summary</h2>
          <div className="mt-6 flex justify-between text-lg">
            <span>Total</span>
            <span className="font-semibold">{formatPrice(subtotal)}</span>
          </div>

          {error && (
            <div className="mt-4 rounded-xl border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-700">
              {error}
            </div>
          )}

          <button
            type="button"
            onClick={handleCheckout}
            disabled={loading || !accountId}
            className="mt-6 flex w-full items-center justify-center gap-2 rounded-full bg-ink py-3.5 text-sm font-semibold text-white transition hover:bg-slate-800 disabled:cursor-not-allowed disabled:opacity-50"
          >
            {loading ? (
              <>
                <Loader2 className="h-4 w-4 animate-spin" />
                Processing...
              </>
            ) : (
              "Place order"
            )}
          </button>

          <button
            type="button"
            onClick={() => navigate("/cart")}
            className="mt-3 w-full text-center text-sm text-muted transition hover:text-ink"
          >
            Back to cart
          </button>
        </div>
      </div>
    </div>
  );
}
