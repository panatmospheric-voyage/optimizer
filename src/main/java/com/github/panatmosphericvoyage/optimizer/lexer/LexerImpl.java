package com.github.panatmosphericvoyage.optimizer.lexer;

import java.io.IOException;
import java.io.InputStream;
import java.util.ArrayList;
import java.util.List;

public class LexerImpl extends Lexer {

	public static final char STATEMENT_TERMINATOR = ';';
	public static final char OPERATOR_EQUALS = '=';
	public static final char OPERATOR_ADD = '+';
	public static final char OPERATOR_SUBTRACT = '-';
	public static final char OPERATOR_MULTIPLY = '*';
	public static final char OPERATOR_DIVIDE = '/';

	private TokenType tokenType = null;
	private List<Token> tokens = null;

	public LexerImpl(InputStream in) {
		super(in);
	}

	@Override
	public List<Token> tokenize() throws IOException, LexerException {
		tokenType = null;
		tokens = new ArrayList<Token>();
		StringBuilder token = new StringBuilder();
		while (hasNext()) {
			char c = read();
			// System.err.print(c);
			if ((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_') {
				// could be an identifier or keyword
				if (tokenType == null) {
					tokenType = TokenType.IDENTIFIER;
					token.append(c);
				} else if (tokenType == TokenType.IDENTIFIER) {
					token.append(c);
				} else if (tokenType == TokenType.NUMBER) {
					if (c == 'e' || c == 'E') {
						token.append(c);
					} else {
						// this must be unit information
						addToken(token);
						tokenType = TokenType.IDENTIFIER;
						token.append(c);
					}
				} else {
					throw new LexerException("Don't know what happens here");
				}
			} else if ((c >= '0' && c <= '9') || c == '.') {
				if (tokenType == null) {
					tokenType = TokenType.NUMBER;
				}
				token.append(c);
			} else if (c == '-') { // could be a negative number or subtraction

			} else if (c == '=') {
				addToken(token);
				tokens.add(new Token(TokenType.OPERATOR_EQUALS, "="));
			} else if (c == '+') {
				addToken(token);
				tokens.add(new Token(TokenType.OPERATOR_ADD, "+"));
			} else if (c == '*') {
				addToken(token);
				tokens.add(new Token(TokenType.OPERATOR_MULTIPLY, "*"));
			} else if (c == '/') {
				if (peek() == '/') {
					// comment
					while (read() != '\n')
						;
				} else {
					addToken(token);
					tokens.add(new Token(TokenType.OPERATOR_DIVIDE, "/"));
				}
			} else if (c == '^') {
				addToken(token);
				tokens.add(new Token(TokenType.OPERATOR_EXPONENT, "^"));
			} else if (c == '{') {
				addToken(token);
				tokens.add(new Token(TokenType.BRACE_OPEN, "{"));
			} else if (c == '}') {
				addToken(token);
				tokens.add(new Token(TokenType.BRACE_CLOSE, "}"));
			} else if (c == '(') {
				addToken(token);
				tokens.add(new Token(TokenType.PARENTHESIS_OPEN, "("));
			} else if (c == ')') {
				addToken(token);
				tokens.add(new Token(TokenType.PARENTHESIS_CLOSE, ")"));
			} else if (c == '[') {
				addToken(token);
				tokens.add(new Token(TokenType.BRACKET_OPEN, "["));
			} else if (c == ']') {
				addToken(token);
				tokens.add(new Token(TokenType.BRACKET_CLOSE, "]"));
			} else if (c == ',') {
				if (tokenType == TokenType.NUMBER) { // handle comma in number
					token.append(c);
				} else {
					addToken(token);
					tokens.add(new Token(TokenType.STATEMENT_SEPARATOR, ","));
				}
			} else if (c == STATEMENT_TERMINATOR) {
				addToken(token);
				tokens.add(new Token(TokenType.STATEMENT_TERMINATOR, String.valueOf(STATEMENT_TERMINATOR)));
			} else {
				addToken(token);
			}
		}
		return tokens;
	}

	public void addToken(StringBuilder token) {
		if (tokenType != null) {
			tokens.add(new Token(tokenType, token.toString()));
			token.setLength(0); // clear token
			tokenType = null; // reset token type
		}
	}

	public String nextIdentifier() throws IOException {
		return nextToken(this::isNotWhitespace, this::isNotIdentifier);
	}

	private boolean isNotWhitespace(char c) {
		return !Character.isWhitespace(c);
	}

	private boolean isNotIdentifier(char c) {
		return !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_');
	}

}
