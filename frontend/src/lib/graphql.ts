const GRAPHQL_URL = import.meta.env.VITE_GRAPHQL_URL ?? "/graphql";

interface GraphQLResponse<T> {
  data?: T;
  errors?: Array<{ message: string }>;
}

export async function graphqlRequest<T>(
  query: string,
  variables?: Record<string, unknown>,
): Promise<T> {
  const response = await fetch(GRAPHQL_URL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ query, variables }),
  });

  if (!response.ok) {
    throw new Error(`GraphQL request failed (${response.status})`);
  }

  const result: GraphQLResponse<T> = await response.json();

  if (result.errors?.length) {
    throw new Error(result.errors.map((e) => e.message).join(", "));
  }

  if (!result.data) {
    throw new Error("No data returned from GraphQL");
  }

  return result.data;
}

export const queries = {
  products: `
    query Products($pagination: PaginationInput, $query: String, $id: String) {
      products(pagination: $pagination, query: $query, id: $id) {
        id
        name
        description
        price
      }
    }
  `,
  accountsList: `
    query AccountsList($pagination: PaginationInput) {
      accounts(pagination: $pagination) {
        id
        name
      }
    }
  `,
  accounts: `
    query Accounts($pagination: PaginationInput, $id: String) {
      accounts(pagination: $pagination, id: $id) {
        id
        name
        orders {
          id
          createdAt
          totalPrice
          products {
            id
            name
            description
            price
            quantity
          }
        }
      }
    }
  `,
  createAccount: `
    mutation CreateAccount($account: AccountInput!) {
      createAccount(account: $account) {
        id
        name
      }
    }
  `,
  createProduct: `
    mutation CreateProduct($product: ProductInput!) {
      createProduct(product: $product) {
        id
        name
        description
        price
      }
    }
  `,
  createOrder: `
    mutation CreateOrder($order: OrderInput!) {
      createOrder(order: $order) {
        id
        createdAt
        totalPrice
        products {
          id
          name
          description
          price
          quantity
        }
      }
    }
  `,
};
