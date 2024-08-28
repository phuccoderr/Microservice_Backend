package com.phuc.reviewservice.models;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Table;
import lombok.Getter;
import lombok.Setter;

@Entity
@Setter
@Getter
@Table(name = "reviews")
public class Review extends BaseEntity {
    private String headline;
    private String comment;
    private int rating;

    @Column(name = "product_id")
    private String productId;

    @Column(name = "customer_id")
    private String customerId;
}
