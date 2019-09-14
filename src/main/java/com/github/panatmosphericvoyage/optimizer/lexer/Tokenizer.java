package com.github.panatmosphericvoyage.optimizer.lexer;

import java.io.IOException;
import java.io.InputStream;
import java.util.function.Function;

import parser.ParserException;

public class Tokenizer {

	private final InputStream in;
	private final Function<Character, Boolean> identifierFilter;
	private char lastChar;
	private int curLine = 0;
	private int curCol = 0;
	private boolean hasLastChar = false;
	private char terminationChar;

	public Tokenizer(InputStream in, Function<Character, Boolean> identifierFilter) {
		this.in = in;
		this.identifierFilter = identifierFilter;
	}

	public boolean hasNext() throws IOException {
		return in.available() > 0;
	}

	public void discardWhitespace() throws IOException {
		while (Character.isWhitespace(read())) {
			// do nothing
		}
		hasLastChar = true;
	}

	public String readIdentifier(char... terminators) throws IOException, ParserException {
		StringBuilder sb = new StringBuilder();

		// read extra character if needed
		if (hasLastChar) {
			sb.append(lastChar);
			hasLastChar = false;
		}

		// read identifier chars
		while (identifierFilter.apply(terminationChar = read()) && !contains(terminators, getLastChar())) {
			sb.append(getLastChar());
		}
		hasLastChar = true;

		return sb.toString();
	}

	public char getTerminationChar() {
		return terminationChar;
	}

	private boolean contains(char[] cs, char c) {
		for (char d : cs) {
			if (d == c) {
				return true;
			}
		}
		return false;
	}

	public char read() throws IOException {
		lastChar = (char) in.read();
		if (lastChar == '\n') {
			curLine++;
			curCol = 0;
		} else {
			curCol++;
		}
		return lastChar;
	}

	public char getLastChar() {
		return lastChar;
	}

	public int getCurrentLine() {
		return curLine;
	}

	public int getCurrentColumn() {
		return curCol;
	}

}
