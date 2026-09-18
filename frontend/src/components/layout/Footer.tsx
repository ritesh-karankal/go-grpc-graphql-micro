import { Link } from "react-router-dom";
import { ShoppingBag } from "lucide-react";

export function Footer() {
  return (
    <footer className="mt-auto border-t border-slate-200 bg-white">
      <div className="mx-auto max-w-7xl px-4 py-12 sm:px-6 lg:px-8">
        <div className="grid gap-8 md:grid-cols-4">
          <div className="md:col-span-2">
            <Link to="/" className="flex items-center gap-2">
              <div className="flex h-9 w-9 items-center justify-center rounded-xl bg-ink text-white">
                <ShoppingBag className="h-4 w-4" />
              </div>
              <span className="font-serif text-2xl">Meridian</span>
            </Link>
            <p className="mt-4 max-w-sm text-sm leading-relaxed text-muted">
              A modern storefront powered by Go microservices, gRPC, and GraphQL.
              Browse curated products, manage your account, and checkout seamlessly.
            </p>
          </div>

          <div>
            <h3 className="text-sm font-semibold uppercase tracking-wider text-ink">
              Shop
            </h3>
            <ul className="mt-4 space-y-2 text-sm text-muted">
              <li>
                <Link to="/shop" className="hover:text-ink transition-colors">
                  All Products
                </Link>
              </li>
              <li>
                <Link to="/cart" className="hover:text-ink transition-colors">
                  Cart
                </Link>
              </li>
              <li>
                <Link to="/orders" className="hover:text-ink transition-colors">
                  Order History
                </Link>
              </li>
            </ul>
          </div>

          <div>
            <h3 className="text-sm font-semibold uppercase tracking-wider text-ink">
              Account
            </h3>
            <ul className="mt-4 space-y-2 text-sm text-muted">
              <li>
                <Link to="/account" className="hover:text-ink transition-colors">
                  Sign In / Register
                </Link>
              </li>
            </ul>
          </div>
        </div>

        <div className="mt-10 flex flex-col gap-2 border-t border-slate-200 pt-8 text-sm text-muted sm:flex-row sm:items-center sm:justify-between">
          <p>&copy; {new Date().getFullYear()} Meridian. All rights reserved.</p>
          <p>GraphQL API at localhost:8000</p>
        </div>
      </div>
    </footer>
  );
}
