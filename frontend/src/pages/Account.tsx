import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { Loader2, LogOut, User, UserPlus } from "lucide-react";
import { useCart } from "../lib/cart";
import { graphqlRequest, queries } from "../lib/graphql";
import type { Account as AccountType } from "../types";

export function Account() {
  const { accountId, accountName, setAccount, clearAccount } = useCart();
  const [name, setName] = useState("");
  const [accounts, setAccounts] = useState<AccountType[]>([]);
  const [loading, setLoading] = useState(false);
  const [fetching, setFetching] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    graphqlRequest<{ accounts: AccountType[] }>(queries.accountsList, {
      pagination: { skip: 0, take: 20 },
    })
      .then((data) => setAccounts(data.accounts))
      .catch((err: Error) => setError(err.message))
      .finally(() => setFetching(false));
  }, []);

  async function handleCreate(event: React.FormEvent) {
    event.preventDefault();
    const trimmed = name.trim();
    if (!trimmed) return;

    setLoading(true);
    setError(null);

    try {
      const data = await graphqlRequest<{ createAccount: AccountType }>(
        queries.createAccount,
        { account: { name: trimmed } },
      );
      setAccount(data.createAccount.id, data.createAccount.name);
      setName("");
      setAccounts((prev) => [data.createAccount, ...prev]);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Could not create account");
    } finally {
      setLoading(false);
    }
  }

  function handleSelect(account: AccountType) {
    setAccount(account.id, account.name);
  }

  return (
    <div className="mx-auto max-w-2xl px-4 py-10 sm:px-6 lg:px-8">
      <h1 className="font-serif text-4xl tracking-tight">Account</h1>
      <p className="mt-2 text-muted">
        Create a new account or select an existing one to checkout and view orders.
      </p>

      {accountId && (
        <div className="mt-8 flex items-center justify-between rounded-2xl border border-brand-200 bg-brand-50 p-5">
          <div className="flex items-center gap-3">
            <div className="flex h-10 w-10 items-center justify-center rounded-full bg-brand-200">
              <User className="h-5 w-5 text-brand-700" />
            </div>
            <div>
              <p className="text-sm text-brand-700">Currently signed in</p>
              <p className="font-semibold">{accountName}</p>
            </div>
          </div>
          <button
            type="button"
            onClick={clearAccount}
            className="inline-flex items-center gap-1.5 rounded-full border border-brand-300 px-4 py-2 text-sm font-medium text-brand-800 transition hover:bg-brand-100"
          >
            <LogOut className="h-4 w-4" />
            Sign out
          </button>
        </div>
      )}

      <form onSubmit={handleCreate} className="mt-10">
        <h2 className="flex items-center gap-2 text-lg font-semibold">
          <UserPlus className="h-5 w-5" />
          Create account
        </h2>
        <div className="mt-4 flex gap-3">
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Your name"
            className="flex-1 rounded-xl border border-slate-200 px-4 py-3 text-sm outline-none transition focus:border-brand-400 focus:ring-2 focus:ring-brand-100"
          />
          <button
            type="submit"
            disabled={loading || !name.trim()}
            className="inline-flex items-center gap-2 rounded-xl bg-ink px-5 py-3 text-sm font-semibold text-white transition hover:bg-slate-800 disabled:opacity-50"
          >
            {loading && <Loader2 className="h-4 w-4 animate-spin" />}
            Create
          </button>
        </div>
      </form>

      {error && (
        <div className="mt-4 rounded-xl border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-700">
          {error}
        </div>
      )}

      <div className="mt-10">
        <h2 className="text-lg font-semibold">Existing accounts</h2>
        {fetching && (
          <div className="mt-4 flex justify-center py-8">
            <Loader2 className="h-6 w-6 animate-spin text-muted" />
          </div>
        )}
        {!fetching && accounts.length === 0 && (
          <p className="mt-4 text-sm text-muted">No accounts yet. Create one above.</p>
        )}
        <ul className="mt-4 space-y-2">
          {accounts.map((account) => (
            <li key={account.id}>
              <button
                type="button"
                onClick={() => handleSelect(account)}
                className={`flex w-full items-center justify-between rounded-xl border px-4 py-3 text-left text-sm transition ${
                  accountId === account.id
                    ? "border-brand-400 bg-brand-50"
                    : "border-slate-200 bg-white hover:border-slate-300"
                }`}
              >
                <span className="font-medium">{account.name}</span>
                {accountId === account.id && (
                  <span className="text-xs font-medium text-brand-600">Active</span>
                )}
              </button>
            </li>
          ))}
        </ul>
      </div>

      {accountId && (
        <Link
          to="/orders"
          className="mt-8 inline-block text-sm font-medium text-brand-600 hover:underline"
        >
          View your order history &rarr;
        </Link>
      )}
    </div>
  );
}
