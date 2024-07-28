package com.phuc.categoryservice.configurations;

import com.phuc.categoryservice.dtos.CategoryDto;
import com.phuc.categoryservice.models.Category;
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

        modelMapper.addMappings(new PropertyMap<CategoryDto,Category>() {
            @Override
            protected void configure() {
                skip(destination.getParent());
                skip(destination.getChildren());
            }
        });

        return modelMapper;
    }
}
