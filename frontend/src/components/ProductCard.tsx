import { Link } from "react-router-dom";
import { ArrowUpRight, ShoppingBag } from "lucide-react";
import type { Product } from "../types";
import { useCart, formatPrice } from "../lib/cart";
import { cn, productImage, productImageUrl } from "../lib/utils";

interface ProductCardProps {
  product: Product;
  featured?: boolean;
}

export function ProductCard({ product, featured = false }: ProductCardProps) {
  const { addItem } = useCart();
  const gradient = productImage(product.name);
  const imageUrl = productImageUrl(product.name);

  return (
    <article
      className={cn(
        "group relative flex flex-col overflow-hidden rounded-3xl border border-slate-200 bg-white transition duration-300 hover:-translate-y-1 hover:shadow-xl hover:shadow-slate-200/60",
        featured && "md:col-span-2 md:flex-row",
      )}
    >
      <Link
        to={`/product/${product.id}`}
        className={cn(
          "relative flex aspect-[4/5] items-center justify-center overflow-hidden bg-gradient-to-br",
          gradient,
          featured && "md:aspect-auto md:w-1/2 md:min-h-[320px]",
        )}
      >
        {imageUrl && (
          <img
            src={imageUrl}
            alt={product.name}
            className="absolute inset-0 h-full w-full object-cover transition duration-500 group-hover:scale-105"
          />
        )}
        <div className="absolute inset-0 bg-[radial-gradient(circle_at_30%_20%,rgba(255,255,255,0.22),transparent_50%)]" />
        {!imageUrl && (
          <span className="relative font-serif text-5xl text-ink/20 transition group-hover:scale-110">
            {product.name.charAt(0).toUpperCase()}
          </span>
        )}
        <div className="absolute right-4 top-4 flex h-10 w-10 items-center justify-center rounded-full bg-white/80 opacity-0 backdrop-blur transition group-hover:opacity-100">
          <ArrowUpRight className="h-4 w-4" />
        </div>
      </Link>

      <div className={cn("flex flex-1 flex-col p-5", featured && "md:justify-center md:p-8")}>
        <div className="flex-1">
          <Link to={`/product/${product.id}`}>
            <h3 className="text-lg font-semibold tracking-tight transition hover:text-brand-600">
              {product.name}
            </h3>
          </Link>
          <p className="mt-2 line-clamp-2 text-sm leading-relaxed text-muted">
            {product.description}
          </p>
        </div>

        <div className="mt-5 flex items-center justify-between gap-3">
          <span className="text-xl font-semibold">{formatPrice(product.price)}</span>
          <button
            type="button"
            onClick={() => addItem(product)}
            className="inline-flex items-center gap-2 rounded-full bg-ink px-4 py-2.5 text-sm font-medium text-white transition hover:bg-slate-800 active:scale-[0.98]"
          >
            <ShoppingBag className="h-4 w-4" />
            Add
          </button>
        </div>
      </div>
    </article>
  );
}
