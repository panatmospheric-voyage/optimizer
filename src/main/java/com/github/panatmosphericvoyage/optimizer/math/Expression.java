package com.github.panatmosphericvoyage.optimizer.math;

import java.util.Map;

import com.github.panatmosphericvoyage.optimizer.model.Identifier;

public abstract class Expression {

	public abstract double evaluate(Map<Identifier, Expression> values);

	public abstract Expression simplify(Map<Identifier, Expression> values);

}
