package com.ybb.aoc;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.List;

public final class Utils {
    private Utils(){}
    public static List<String> readResource(String res) throws IOException {
        var file = Utils.class.getResource(res);
        List<String> lines = null;
        try(InputStream stream = file.openStream()) {
            lines = new BufferedReader(new InputStreamReader(stream)).lines().toList();
        } catch (IOException e) {
            System.out.println(e.toString());
            throw e;
        }
        return lines;
    }

    public static List<String> readFile(String path) throws Exception {
        List<String> lines = null;
        try(var stream = new FileInputStream(new File(path))) {
            lines = new BufferedReader(new InputStreamReader(stream)).lines().toList();
        } catch (Exception e) {
            System.out.println(e.toString());
            throw e;
        } 
        return lines;
    }
}
