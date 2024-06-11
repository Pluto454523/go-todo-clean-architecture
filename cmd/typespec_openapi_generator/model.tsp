
model OrderListResponse {
    items: OrderResponse[];
    total: int32;
  }
  
  model OrderRequest {
    id: string;
    order_number: string;
    store_id: string;
    provider_code: string;
    flow: string;
    state: string;
    items: OrderItem[];
  }
  
  model OrderResponse {
    id: string;
    order_number: string;
    store_id: string;
    provider_code: string;
    flow: string;
    state: string;
    items: OrderItemWithProduct[];
  }
  
  model OrderItem {
    price: float64;
    qty: int32;
    product_id: string;
  }
  
  model OrderItemWithProduct {
    price: float64;
    qty: int32;
    product: Product;
  }
  
  model Product {
    id: string;
    name: string;
    price: float64;
    images: string[];
  }
  
  model CreateStoreRequest {
    provider_code: string;
    company_id: string;
    name: string;
    code: string;
    address: Address;
    tax_information: TaxInformation;
  }
  
  model StoreModel {
    id: string;
    provider_code: string;
    company_id: string;
    name: string;
    code: string;
    address: Address;
    tax_information: TaxInformation;
  }
  
  model Address {
    street_address: string;
    sub_district: string;
    district: string;
    province: string;
    postal_code: string;
    country: string;
  }
  
  model TaxInformation {
    full_name: string;
    tax_id: string;
    address: Address;
  }
  
  model ErrorResponse {
    error: string;
    error_code: string;
    issue_id: string;
  }
  
  
  
  