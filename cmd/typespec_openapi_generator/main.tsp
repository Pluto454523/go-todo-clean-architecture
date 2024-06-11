import "@typespec/http";
import "@typespec/versioning";
import "@typespec/openapi3";
import "./model.tsp";

using TypeSpec.Http;
using TypeSpec.Versioning;

@service({
  title: "Demo API",
})
@doc("Demo API for the OpenAPI documentation.")
@useAuth(ApiKeyAuth<ApiKeyLocation.header, "X-USER-ID">)
namespace DemoApi;

@route("/company")
namespace Company {
  @route("/{company_id}/store")
  namespace CompanyStore {
    @doc("Create a new store, to be managed by the company.")
    @post op createStore(
      @path company_id: string,
      @body request: CreateStoreRequest,
    ): {
      @statusCode _: 201 | 401 | 500;
      @body body: string | ErrorResponse;
    };

    @doc("List all stores managed by the company.")
    @get op listStores(@path company_id: string): {
      @statusCode _: 200 | 500;
      @body body: StoreModel[] | ErrorResponse;
    };
  }
}

@route("/store")
namespace Store {
  @doc("Get a list of stores that the user has access to.")
  @get op listStores(): {
    @statusCode _: 200 | 500;
    @body body: StoreModel[] | ErrorResponse;
  };


  @route("/{store_id}/order")
  namespace StoreOrder {
    @post op createOrder(
      @path store_id: string,
      @body request: OrderRequest,
    ): {
      @statusCode _: 201 | 401 | 500;
      @body body: string | ErrorResponse;
    };

    @get op listOrders(
      @path store_id: string,
      @query page: int32,
      @query limit: int32,
    ): {
      @statusCode _: 200 | 500;
      @body body: OrderListResponse | ErrorResponse;
    };
  }
}

@route("/order")
namespace Order {
  @doc("Get a list of orders, with dynamic filtering.")
  @get op getOrder(
    @query product_id?: string,
    @query store_id?: string,
    @query limit: int64,
    @query offset: int64,
  ): {
    @statusCode _: 200 | 404 | 500;
    @body body: OrderResponse | ErrorResponse;
  };
}
