package com.phuc.categoryservice.util;

import com.phuc.categoryservice.dtos.CategoryDto;
import com.phuc.categoryservice.models.BaseEntity;
import com.phuc.categoryservice.models.Category;
import org.modelmapper.ModelMapper;

import java.text.Normalizer;
import java.util.regex.Pattern;

public class Utility {
    static ModelMapper modelMapper = new ModelMapper();
    public static String unAccent(String s) {
        String normalizer = Normalizer.normalize(s,Normalizer.Form.NFD);
        Pattern pattern = Pattern.compile("\\p{InCombiningDiacriticalMarks}+");
        String noAccents = pattern.matcher(normalizer).replaceAll("");
        return noAccents.toLowerCase().replace(" ", "-");
    }

    public static CategoryDto toDto(Category category) {
        CategoryDto dto = modelMapper.map(category, CategoryDto.class);

        dto.setChildren(category.getChildren().stream().map(BaseEntity::getId).toList());
        if (category.getParent() != null) {
            dto.setParent(category.getParent().getId());
        }

        return dto;
    }

    public static Category toEntity(CategoryDto dto) {
        return modelMapper.map(dto, Category.class);
    }
}
