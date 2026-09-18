export interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
}

export interface Account {
  id: string;
  name: string;
  orders: Order[];
}

export interface OrderedProduct {
  id: string;
  name: string;
  description: string;
  price: number;
  quantity: number;
}

export interface Order {
  id: string;
  createdAt: string;
  totalPrice: number;
  products: OrderedProduct[];
}

export interface CartItem {
  product: Product;
  quantity: number;
}

export interface PaginationInput {
  skip?: number;
  take?: number;
}
