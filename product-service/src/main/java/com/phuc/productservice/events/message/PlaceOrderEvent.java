package com.phuc.productservice.events.message;


import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Getter;
import lombok.Setter;

import java.util.List;

@Setter
@Getter
public class PlaceOrderEvent {

    @JsonProperty("customer_id")
    private String customerId;

    @JsonProperty("customer_email")
    private String customerEmail;

    @JsonProperty("items")
    private List<Item> items;

    @JsonProperty("address")
    private String address;

    @JsonProperty("payment_method")
    private String paymentMethod;

    @JsonProperty("check_out")
    private CheckOut checkOut;

    // Getters and setters...

    @Setter
    @Getter
    public static class Item {
        @JsonProperty("product_id")
        private String productId;

        @JsonProperty("quantity")
        private int quantity;
    }


    @Setter
    @Getter
    public static class CheckOut {
        @JsonProperty("product_total")
        private Float productTotal;

        @JsonProperty("product_cost")
        private Float productCost;

        @JsonProperty("shipping_cost")
        private Float shippingCost;

        @JsonProperty("deliver_days")
        private String deliverDays;
    }
}
