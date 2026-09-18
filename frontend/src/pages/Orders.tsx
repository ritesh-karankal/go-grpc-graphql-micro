import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { ClipboardList, Loader2 } from "lucide-react";
import { useCart, formatPrice } from "../lib/cart";
import { graphqlRequest, queries } from "../lib/graphql";
import { formatDate, productImage, productImageUrl } from "../lib/utils";
import { EmptyState } from "../components/EmptyState";
import type { Account, Order } from "../types";

export function Orders() {
  const { accountId, accountName } = useCart();
  const [orders, setOrders] = useState<Order[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!accountId) return;

    setLoading(true);
    setError(null);

    graphqlRequest<{ accounts: Account[] }>(queries.accounts, { id: accountId })
      .then((data) => setOrders(data.accounts[0]?.orders ?? []))
      .catch((err: Error) => setError(err.message))
      .finally(() => setLoading(false));
  }, [accountId]);

  if (!accountId) {
    return (
      <div className="mx-auto max-w-7xl px-4 py-10 sm:px-6 lg:px-8">
        <EmptyState
          icon={ClipboardList}
          title="Sign in to view orders"
          description="Create or select an account to see your order history."
          actionLabel="Go to account"
          actionTo="/account"
        />
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-7xl px-4 py-10 sm:px-6 lg:px-8">
      <h1 className="font-serif text-4xl tracking-tight">Order history</h1>
      <p className="mt-2 text-muted">Orders for {accountName}</p>

      {loading && (
        <div className="flex justify-center py-20">
          <Loader2 className="h-8 w-8 animate-spin text-muted" />
        </div>
      )}

      {error && (
        <div className="mt-8 rounded-2xl border border-red-200 bg-red-50 px-6 py-4 text-sm text-red-700">
          {error}
        </div>
      )}

      {!loading && !error && orders.length === 0 && (
        <div className="mt-8">
          <EmptyState
            icon={ClipboardList}
            title="No orders yet"
            description="When you place an order, it will appear here."
            actionLabel="Start shopping"
            actionTo="/shop"
          />
        </div>
      )}

      {!loading && orders.length > 0 && (
        <div className="mt-10 space-y-4">
          {orders.map((order) => (
            <div
              key={order.id}
              className="rounded-2xl border border-slate-200 bg-white p-6"
            >
              <div className="flex flex-wrap items-start justify-between gap-4">
                <div>
                  <p className="font-mono text-xs text-muted">{order.id}</p>
                  <p className="mt-1 text-sm text-muted">{formatDate(order.createdAt)}</p>
                </div>
                <p className="text-xl font-semibold">{formatPrice(order.totalPrice)}</p>
              </div>

              <ul className="mt-5 divide-y divide-slate-100">
                {order.products.map((product) => (
                  <li
                    key={`${order.id}-${product.id}`}
                    className="flex items-center justify-between gap-4 py-3 text-sm"
                  >
                    <div className="flex min-w-0 items-center gap-3">
                      <div
                        className={`relative flex h-12 w-12 shrink-0 items-center justify-center overflow-hidden rounded-lg bg-gradient-to-br ${productImage(product.name)}`}
                      >
                        {productImageUrl(product.name) ? (
                          <img
                            src={productImageUrl(product.name)}
                            alt={product.name}
                            className="h-full w-full object-cover"
                          />
                        ) : (
                          <span className="font-serif text-lg text-ink/20">
                            {product.name.charAt(0).toUpperCase()}
                          </span>
                        )}
                      </div>
                      <span className="truncate">
                        {product.name}{" "}
                        <span className="text-muted">&times; {product.quantity}</span>
                      </span>
                    </div>
                    <span className="font-medium">
                      {formatPrice(product.price * product.quantity)}
                    </span>
                  </li>
                ))}
              </ul>
            </div>
          ))}
        </div>
      )}

      <Link
        to="/shop"
        className="mt-8 inline-block text-sm font-medium text-brand-600 hover:underline"
      >
        Continue shopping &rarr;
      </Link>
    </div>
  );
}
