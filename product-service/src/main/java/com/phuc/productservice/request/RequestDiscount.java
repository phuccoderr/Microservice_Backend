package com.phuc.productservice.request;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;
import lombok.Getter;
import lombok.Setter;

import java.time.LocalDateTime;

@Setter
@Getter
public class RequestDiscount {

    @JsonProperty("name")
    @NotBlank(message = "name cannot empty or null")
    private String name;

    @JsonProperty("code")
    @NotBlank(message = "code cannot empty or null")
    private String code;

    @JsonProperty("expiry_date")
    private LocalDateTime expiryDate;

    @JsonProperty("sale")
    @Min(value = 0, message = "sale cannot be less than 0")
    @Max(value = 100, message = "sale cannot greater than 100")
    private Integer sale;

    @JsonProperty("quantity")
    @Min(value = 0, message = "sale cannot be less than 0")
    private Integer quantity;
}
