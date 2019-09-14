package parser;

import java.io.IOException;
import java.io.InputStream;

import com.github.panatmosphericvoyage.optimizer.lexer.Lexer;
import com.github.panatmosphericvoyage.optimizer.model.OptimizerModel;

public abstract class Parser {

	public static final String KEYWORD_UNIT = "unit";
	
	
	protected final Lexer lexer;

	public Parser(Lexer lexer) {
		this.lexer = lexer;
	}

	public abstract OptimizerModel parse(InputStream in) throws IOException, ParserException;

	public OptimizerModel parse(String s) throws ParserException {
		try {
			return parse(new InputStream() {
				int index = 0;

				@Override
				public int read() throws IOException {
					return s.charAt(index++);
				}

			});
		} catch (IOException e) {
			e.printStackTrace();
		}
		return null;
	}

}
