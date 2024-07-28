package com.phuc.categoryservice.util;

import com.phuc.categoryservice.dtos.CategoryDto;
import com.phuc.categoryservice.exceptions.ParamValidateException;
import com.phuc.categoryservice.models.BaseEntity;
import com.phuc.categoryservice.models.Category;
import org.modelmapper.ModelMapper;

import java.text.Normalizer;
import java.util.List;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

public class Utility {
    static ModelMapper modelMapper = new ModelMapper();
    public static String unAccent(String s) {
        String normalizer = Normalizer.normalize(s,Normalizer.Form.NFD);
        Pattern pattern = Pattern.compile("\\p{InCombiningDiacriticalMarks}+");
        String noAccents = pattern.matcher(normalizer).replaceAll("");
        return noAccents.toLowerCase().replace(" ", "-");
    }

    public static List<CategoryDto> toListDtos(List<Category> categories) {
        return categories.stream().map(Utility::toDto).toList();
    }

    public static CategoryDto toDto(Category category) {
        CategoryDto dto = modelMapper.map(category, CategoryDto.class);

        dto.setChildren(category.getChildren().stream().map(BaseEntity::getId).toList());
        if (category.getParent() != null) {
            dto.setParent(category.getParent().getId());
        }
        return dto;
    }

    public static void checkSortIsAscOrDesc(String sort) throws ParamValidateException {
        if (!sort.equals("asc") && !sort.equals("desc")) {
            throw new ParamValidateException();
        }
    }

}
