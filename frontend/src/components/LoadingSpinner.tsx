import { cn } from "../lib/utils";

export function LoadingSpinner({ className }: { className?: string }) {
  return (
    <div className={cn("flex items-center justify-center py-20", className)}>
      <div className="h-10 w-10 animate-spin rounded-full border-2 border-slate-200 border-t-brand-500" />
    </div>
  );
}
