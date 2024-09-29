package com.phuc.categoryservice.dtos;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.phuc.categoryservice.models.Category;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;

@Setter
@Getter
@NoArgsConstructor
@AllArgsConstructor
public class CategoryDto {

    @JsonProperty("id")
    protected String id;

    @JsonProperty("created_at")
    private LocalDateTime createdAt;

    @JsonProperty("updated_at")
    private LocalDateTime updatedAt;

    @JsonProperty("name")
    private String name;

    @JsonProperty("alias")
    private String alias;

    @JsonProperty("status")
    private boolean status;

    @JsonProperty("has_children")
    private boolean hasChildren;

    @JsonProperty("parent_id")
    private String parentId;

    @JsonProperty("children")
    private List<CategoryDto> children = new ArrayList<>();

    public static CategoryDto copyIdAndName(Category category) {
        CategoryDto copyCategory = new CategoryDto();
        copyCategory.setId(category.getId());
        copyCategory.setName(category.getName());

        return copyCategory;
    }
}
