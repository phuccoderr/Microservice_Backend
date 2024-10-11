package com.phuc.productservice.models;

import jakarta.persistence.*;
import lombok.*;

import java.time.LocalDateTime;

@Setter
@Getter
@AllArgsConstructor
@NoArgsConstructor
@Builder
@Entity
@Table(name = "discounts")
public class Discount extends BaseEntity {

    private String name;

    private String code;

    @Column(name = "expiry_date")
    private LocalDateTime expiryDate;
    private Integer sale;
    private Integer quantity;

}
