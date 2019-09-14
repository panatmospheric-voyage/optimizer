package com.github.panatmosphericvoyage.optimizer;

import java.io.File;
import java.io.FileInputStream;
import java.util.List;

import com.github.panatmosphericvoyage.optimizer.lexer.LexerImpl;
import com.github.panatmosphericvoyage.optimizer.lexer.Token;

public class Main {

	public static void main(String[] args) {

		try {
			File file = new File("sample_program.txt");
			System.out.println("File path: " + file.getAbsolutePath());
			LexerImpl lexer = new LexerImpl(new FileInputStream(file));
			List<Token> tokens = lexer.tokenize();
			for (Token t : tokens) {
				System.out.printf("TokenType: %s, Value: %s\n", t.getType().name(), t.getValue());
			}
		} catch (Exception e) {
			e.printStackTrace();
		}

	}

}
