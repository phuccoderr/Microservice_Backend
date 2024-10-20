package com.phuc.productservice.events.message;


import com.fasterxml.jackson.annotation.JsonProperty;
import com.phuc.productservice.dtos.ProductDto;
import lombok.Getter;
import lombok.Setter;

import java.util.List;

@Setter
@Getter
public class PlaceOrderEvent {

    @JsonProperty("customer_id")
    private Customer customerId;

    @Setter
    @Getter
    public static class Customer {
        @JsonProperty("_id")
        private String id;

        @JsonProperty("email")
        private String email;

        @JsonProperty("name")
        private String name;
    }

    @JsonProperty("address")
    private String address;

    @JsonProperty("payment_method")
    private String paymentMethod;

    @JsonProperty("phone_number")
    private String phoneNumber;

    @JsonProperty("total")
    private Float total;

    @JsonProperty("product_cost")
    private Float productCost;

    @JsonProperty("shipping_cost")
    private Float shippingCost;

    @JsonProperty("deliver_days")
    private String deliverDays;


    @JsonProperty("order_details")
    private List<OrderDetail> OrderDetails;

    @Setter
    @Getter
    public static class OrderDetail {
        @JsonProperty("product_id")
        private ProductDto productId;

        @JsonProperty("quantity")
        private int quantity;
    }

}
