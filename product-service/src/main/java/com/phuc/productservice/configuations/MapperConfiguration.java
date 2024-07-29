package com.phuc.productservice.configuations;

import com.phuc.productservice.dtos.ProductDto;
import com.phuc.productservice.models.Product;
import org.modelmapper.ModelMapper;
import org.modelmapper.PropertyMap;
import org.modelmapper.convention.MatchingStrategies;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class MapperConfiguration {

    @Bean
    public ModelMapper modelMapper() {
        ModelMapper modelMapper = new ModelMapper();
        modelMapper.getConfiguration().setMatchingStrategy(MatchingStrategies.STRICT);

        modelMapper.addMappings(new PropertyMap<ProductDto, Product>() {
            @Override
            protected void configure() {
                skip(destination.getExtraImages());
            }
        });

        return modelMapper;
    }
}
