package com.phuc.reviewservice.events.message;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

import java.util.List;

@Setter
@Getter
@Builder
public class PlaceOrderEvent {

    @JsonProperty("customer_id")
    private String customerId;

    @JsonProperty("customer_email")
    private String customerEmail;

    @JsonProperty("customer_name")
    private String customerName;

    @JsonProperty("address")
    private String address;

    @JsonProperty("payment_method")
    private String paymentMethod;

    @JsonProperty("phone_number")
    private String phoneNumber;

    @JsonProperty("check_out")
    private CheckOut checkOut;

    @JsonProperty("items")
    private List<Item> items;

    @Setter
    @Getter
    public static class Item {
        @JsonProperty("product_id")
        private String productId;

        @JsonProperty("product_image")
        private String productImage;

        @JsonProperty("quantity")
        private int quantity;

        @JsonProperty("cost")
        private float cost;

        @JsonProperty("price")
        private float price;

        @JsonProperty("total")
        private float total;
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
