package com.github.panatmosphericvoyage.optimizer.lexer;

public class Token {
	public TokenType type;
	public String value;

	public Token(TokenType type, String value) {
		this.type = type;
		this.value = value;
	}

	public TokenType getType() {
		return type;
	}

	public void setType(TokenType type) {
		this.type = type;
	}

	public String getValue() {
		return value;
	}

	public void setValue(String value) {
		this.value = value;
	}

}
