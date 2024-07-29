package com.phuc.productservice.util;

import com.phuc.productservice.dtos.ProductDto;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.models.ProductImage;
import lombok.experimental.UtilityClass;
import org.modelmapper.ModelMapper;

import java.text.Normalizer;
import java.util.List;
import java.util.regex.Pattern;

@UtilityClass
public class Utility {
    static ModelMapper modelMapper = new ModelMapper();
    public String unAccent(String s) {
        String normalizer = Normalizer.normalize(s,Normalizer.Form.NFD);
        Pattern pattern = Pattern.compile("\\p{InCombiningDiacriticalMarks}+");
        String noAccents = pattern.matcher(normalizer).replaceAll("");
        return noAccents.toLowerCase().replace(" ", "-");
    }

    public ProductDto toDto(Product product) {
        ProductDto dto = modelMapper.map(product, ProductDto.class);

        if (product.getExtraImages() != null) {
            List<String> list = product.getExtraImages().stream().map(ProductImage::getUrl).toList();
            dto.setExtraImages(list);
        }

        return dto;
    }
}
