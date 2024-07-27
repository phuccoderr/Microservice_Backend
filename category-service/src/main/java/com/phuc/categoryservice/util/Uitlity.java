package com.phuc.categoryservice.util;

import java.text.Normalizer;

public class Uitlity {

    public static String unAccent(String s) {
        return Normalizer
                .normalize(s,Normalizer.Form.NFD)
                .replaceAll("[^\\p{ASCII}]", "")
                .replace(" ","-");

    }
}
