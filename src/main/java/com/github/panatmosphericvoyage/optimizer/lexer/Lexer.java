package com.github.panatmosphericvoyage.optimizer.lexer;

import java.io.IOException;
import java.io.InputStream;
import java.util.List;
import java.util.function.Function;

public abstract class Lexer {

	public static final int BUFFER_SIZE = 64;

	private char[] buffer = new char[BUFFER_SIZE];
	int index = 0;
	int endIndex = -1;
	private InputStream in;

	/**
	 * Creates a new {@code Lexer} that reads from the given
	 * {@code InputStream}.
	 * 
	 * @param in the {@code InputStream} to read from
	 */
	public Lexer(InputStream in) {
		this.in = in;
		this.index = buffer.length;
	}

	/**
	 * Converts the given {@code InputStream}'s contents into a list of tokens
	 * for a parser to parse.
	 * 
	 * @return the list of tokens
	 * @throws IOException
	 */
	public abstract List<Token> tokenize() throws IOException, LexerException;

	/**
	 * Reads from the given {@code InputStream} and returns a token. Using the
	 * startFilter, characters are discarded until {@code true} is returned.
	 * Once {@code true} is returned, characters will be included in the
	 * identifier until stopFilter returns {@code true}.
	 * 
	 * @param startFilter the start character filter
	 * @param stopFilter the stop character filter
	 * @return the token
	 * @throws IOException
	 */
	public String nextToken(Function<Character, Boolean> startFilter, Function<Character, Boolean> stopFilter) throws IOException {
		// discard characters
		while (!startFilter.apply(peek())) {
			read();
		}

		// add characters to identifier
		StringBuilder sb = new StringBuilder();
		while (!stopFilter.apply(peek())) {
			sb.append(read());
		}

		return sb.toString();
	}

	/**
	 * Checks if the stream has more characters to read.
	 * 
	 * @return {@code true} if there are more characters, {@code false}
	 *         otherwise
	 */
	public boolean hasNext() {
		return index != endIndex - 1;
	}

	/**
	 * Reads the character the stream is currently at without advancing the
	 * stream.
	 * 
	 * @return the current character
	 * @throws IOException
	 */
	public char peek() throws IOException {
		return buffer[index];
	}

	/**
	 * Reads a character from the stream and advances the stream.
	 * 
	 * @return the character
	 * @throws IOException
	 */
	public char read() throws IOException {
		// fill buffer if at end
		if (index == buffer.length) {
			index = 0;
			fillBuffer();
		}

		// check if we are at the end of the stream
		if (index == endIndex) {
			throw new IOException("End of stream has been reached");
		}

		return buffer[index++];
	}

	private void fillBuffer() throws IOException {
		for (int i = 0; i < buffer.length; i++) {
			int c = in.read();
			if (c == -1) {
				endIndex = i;
				break;
			}
			buffer[i] = (char) c;
		}
	}

}
