package com.phuc.categoryservice.request;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import lombok.Getter;
import lombok.Setter;

import java.util.ArrayList;
import java.util.List;

@Getter
@Setter
public class RequestUpdateCategory {

    @JsonProperty("name")
    @NotBlank(message = "name cannot empty or null")
    private String name;

    @JsonProperty("status")
    @NotNull(message = "status must be true or false")
    private Boolean status;

    @JsonProperty("parent_id")
    @NotNull(message = "parent_id cannot null")
    private String parent;
}
