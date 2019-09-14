package com.github.panatmosphericvoyage.optimizer.math;

import java.util.Map;

import com.github.panatmosphericvoyage.optimizer.model.Identifier;

public class Addition extends Expression {

	private Expression[] addends;

	public Addition(Expression... addends) {
		this.addends = addends;
	}

	@Override
	public double evaluate(Map<Identifier, Expression> values) {
		double sum = 0;
		for (Expression ex : addends) {
			sum += ex.evaluate(values);
		}
		return sum;
	}

	@Override
	public Expression simplify(Map<Identifier, Expression> values) {
		// TODO Auto-generated method stub
		return null;
	}

}
