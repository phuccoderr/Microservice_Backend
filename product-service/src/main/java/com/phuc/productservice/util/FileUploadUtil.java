package com.phuc.productservice.util;

import com.phuc.productservice.constants.Constants;
import com.phuc.productservice.exceptions.FuncErrorException;
import lombok.experimental.UtilityClass;
import org.springframework.web.multipart.MultipartFile;

import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

@UtilityClass
public class FileUploadUtil {

    public static final long MAX_FILE_SIZE = 2 * 1024 * 1024; // 2MB

    public static final String IMAGE_PATTERN = "([^\\s]+(\\.(?i)(jpe?g|png|gif|bmp))$)"; // regex .* file

    public static final String DATE_FORMAT = "yyyyMMddHHmmss";

    public static final String FILE_NAME_FORMAT = "%s_%s";

    public void assertAllowed(MultipartFile file, String pattern) throws FuncErrorException {
//        long size = file.getSize();
//        if (size > MAX_FILE_SIZE) {
//            throw new FuncErrorException(Constants.FILE_MAX_SIZE);
//        }

        String fileName = file.getOriginalFilename();
        if (!isAllowedExtension(fileName,pattern)) {
            throw new FuncErrorException(Constants.FILE_EXTENSION_FAIL);
        }
    }

    public boolean isAllowedExtension(String fileName, String pattern) {
        Matcher matcher = Pattern.compile(pattern, Pattern.CASE_INSENSITIVE).matcher(fileName);
        return matcher.matches();
    }
    public String getFileName(String name) {
        DateFormat dateFormat = new SimpleDateFormat(DATE_FORMAT);
        String date = dateFormat.format(System.currentTimeMillis());
        return String.format(FILE_NAME_FORMAT,name,date);
    }
}
